package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log/slog"
	"votehub-sync/database"
	"votehub-sync/sync"

	_ "turso.tech/database/tursogo"
)

//go:embed sqlc/schema.sql
var dll string

func main() {
	slog.Info("Running votehub sync.")

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
	err = sync.SyncPollTypes(queries)
	if err != nil {
		panic(err)
	}
}
