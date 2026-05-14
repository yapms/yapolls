package votehub

import (
	"context"
	"datasync/database"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type Subjects []Subject

type Subject struct {
	Subject   string    `json:"subject"`
	PollTypes PollTypes `json:"poll_types"`
}

func SyncSubjects(db *database.Queries) error {
	subjects, err := downloadSubjects()
	if err != nil {
		return err
	}

	err = insertSubjects(db, subjects)
	if err != nil {
		return err
	}

	return nil
}

func downloadSubjects() (Subjects, error) {
	response, err := http.Get("https://api.votehub.com/subjects")
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	subjects := Subjects{}
	err = json.Unmarshal(data, &subjects)
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func insertSubjects(db *database.Queries, subjects Subjects) error {
	ctx := context.Background()
	for _, subject := range subjects {
		id, err := uuid.NewV7()
		if err != nil {
			return err
		}

		bid, err := id.MarshalBinary()
		if err != nil {
			return err
		}

		_, err = db.CreateSubject(ctx, database.CreateSubjectParams{
			ID:   bid,
			Name: subject.Subject,
		})
		if err != nil {
			continue
		}

		for _, polltype := range subject.PollTypes {
			polltype, err := db.SearchPollType(ctx, polltype)
			if err != nil {
				continue
			}

			db.CreateSubjectPollType(ctx, database.CreateSubjectPollTypeParams{
				SubjectID:  bid,
				PolltypeID: polltype.ID,
			})
		}
	}

	return nil
}
