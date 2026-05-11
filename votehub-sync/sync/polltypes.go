package sync

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"votehub-sync/database"

	"github.com/google/uuid"
)

type PollTypes []string

func SyncPollTypes(db *database.Queries) error {
	polltypes, err := downloadPollTypes()
	if err != nil {
		return err
	}

	err = insertPollTypes(db, polltypes)
	if err != nil {
		return err
	}

	return nil
}

func downloadPollTypes() (PollTypes, error) {
	response, err := http.Get("https://api.votehub.com/poll-types")
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	polltypes := PollTypes{}
	err = json.Unmarshal(data, &polltypes)
	if err != nil {
		return nil, err
	}

	return polltypes, nil
}

func insertPollTypes(db *database.Queries, pollTypes PollTypes) error {
	ctx := context.Background()
	for _, pollType := range pollTypes {

		id, err := uuid.NewV7()
		if err != nil {
			return err
		}

		bid, err := id.MarshalBinary()
		if err != nil {
			return err
		}

		_, err = db.CreatePollType(ctx, database.CreatePollTypeParams{
			ID:   bid,
			Name: pollType,
		})

		if err != nil {
			continue
		}
	}

	return nil
}

func listPollTypes(db *database.Queries) {
	ctx := context.Background()
	pollTypes, err := db.ListPollTypes(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(pollTypes)
}
