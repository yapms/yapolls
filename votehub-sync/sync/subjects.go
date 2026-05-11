package sync

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"votehub-sync/database"
)

type Subjects []Subject

type Subject struct {
	subject    string
	poll_types PollTypes
}

func SyncSubjects(db *database.Queries) error {
	subjects, err := downloadSubjects()
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
	}
}
