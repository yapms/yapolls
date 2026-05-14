package main

import (
	"context"
	"database/sql"
	"datasync/database"
	"datasync/votehub"
	_ "embed"
	"log/slog"

	_ "turso.tech/database/tursogo"
)

//go:embed sqlc/schema.sql
var dll string

func main() {
	slog.Info("Running votehub votehub.")

	ctx := context.Background()

	slog.Info("Creating database connection.")
	db, err := sql.Open("turso", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, dll)
	if err != nil {
		panic(err)
	}

	queries := database.New(db)

	slog.Info("Syncing poll types.")
	err = votehub.SyncPollTypes(queries)
	if err != nil {
		panic(err)
	}

	slog.Info("Syncing subjects.")
	err = votehub.SyncSubjects(queries)
	if err != nil {
		panic(err)
	}

	slog.Info("Syncing pollsters.")
	err = votehub.SyncPollsters(queries)
	if err != nil {
		panic(err)
	}

	slog.Info("Syncing polls.")
	err = votehub.SyncPolls(queries)
	if err != nil {
		panic(err)
	}
}
