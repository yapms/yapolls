package sync

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"votehub-sync/database"

	"github.com/google/uuid"
)

type Polls []Poll

type Poll struct {
	VotehubID  string   `json:"id"`
	PollType   string   `json:"poll_type"`
	SampleSize int      `json:"sample_size"`
	Population string   `json:"population"`
	URL        string   `json:"url"`
	CreatedAt  string   `json:"created_at"`
	StartDate  string   `json:"start_date"`
	EndDate    string   `json:"end_date"`
	Pollster   string   `json:"pollster"`
	Answers    []Answer `json:"answers"`
	SeatName   string   `json:"seat_name"`
	Sponsors   []string `json:"sponsors"`
	Internal   bool     `json:"internal"`
	Partisan   string   `json:"partisan"`
	Subject    string   `json:"subject"`
}

type Answer struct {
	Choice string  `json:"choice"`
	PCT    float32 `json:"pct"`
}

func SyncPolls(db *database.Queries) error {
	polls, err := downloadPolls()
	if err != nil {
		return err
	}

	err = insertPolls(db, polls)
	if err != nil {
		return err
	}

	return nil
}

func downloadPolls() (Polls, error) {
	response, err := http.Get("https://api.votehub.com/polls")
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	polls := Polls{}
	err = json.Unmarshal(data, &polls)
	if err != nil {
		return nil, err
	}

	return polls, nil
}

func insertPolls(db *database.Queries, polls Polls) error {
	ctx := context.Background()

	for _, poll := range polls {
		id, err := uuid.NewV7()
		if err != nil {
			return err
		}

		bid, err := id.MarshalBinary()
		if err != nil {
			return err
		}

		pollster, err := db.SearchPollster(ctx, poll.Pollster)
		if err != nil {
			continue
		}

		subject, err := db.SearchSubject(ctx, poll.Subject)
		if err != nil {
			continue
		}

		internal := int64(0)
		if poll.Internal {
			internal = 1
		}

		_, err = db.CreatePoll(ctx, database.CreatePollParams{
			ID:        bid,
			VotehubID: poll.VotehubID,
			PollType:  poll.PollType,
			SampleSize: sql.NullInt64{
				Int64: int64(poll.SampleSize),
				Valid: poll.SampleSize != 0,
			},
			Population: sql.NullString{
				String: poll.Population,
				Valid:  poll.Population != "",
			},
			Url:        poll.URL,
			CreatedAt:  poll.CreatedAt,
			StartDate:  poll.StartDate,
			EndDate:    poll.EndDate,
			PollsterID: pollster.ID,
			SeatName: sql.NullString{
				String: poll.SeatName,
				Valid:  poll.SeatName != "",
			},
			Internal: internal,
			Partisan: sql.NullString{
				String: poll.Partisan,
				Valid:  poll.Partisan != "",
			},
			SubjectID: subject.ID,
		})
		if err != nil {
			continue
		}
	}

	return nil
}
