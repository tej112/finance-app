package main

import (
	"context"
	"database/sql"
	"finance/ent"
	"finance/ent/loan"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(100)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	client := Open("host=localhost port=5432 user=finance dbname=finance password=finance sslmode=disable")
	defer client.Close()
	// client = client.Debug()
	err := client.Schema.Create(
		context.Background(),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	)
	if err != nil {
		log.Fatal(err)
	}

}
