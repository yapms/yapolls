package sync

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"votehub-sync/database"

	"github.com/google/uuid"
)

type Pollsters []string

func SyncPollsters(db *database.Queries) error {
	pollsters, err := downloadPollsters()
	if err != nil {
		return err
	}

	err = insertPollsters(db, pollsters)
	if err != nil {
		return err
	}

	return nil
}

func downloadPollsters() (Pollsters, error) {
	response, err := http.Get("https://api.votehub.com/pollsters")
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	pollsters := Pollsters{}
	err = json.Unmarshal(data, &pollsters)
	if err != nil {
		return nil, err
	}

	return pollsters, nil
}

func insertPollsters(db *database.Queries, pollsters Pollsters) error {
	ctx := context.Background()

	for _, pollster := range pollsters {
		id, err := uuid.NewV7()
		if err != nil {
			return err
		}

		bid, err := id.MarshalBinary()
		if err != nil {
			return err
		}

		_, err = db.CreatePollster(ctx, database.CreatePollsterParams{
			ID:   bid,
			Name: pollster,
		})
		if err != nil {
			continue
		}
	}

	return nil
}
