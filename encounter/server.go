package main

import (
	"github.com/smallretardedfish/gql-federation/encounter/faker"
	"github.com/smallretardedfish/gql-federation/encounter/graph"
	"github.com/smallretardedfish/gql-federation/encounter/storage"
	"golang.org/x/exp/slog"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"database/sql"
	_ "github.com/lib/pq"
)

const (
	defaultPort          = "4002"
	defaultDbConnString  = "postgresql://postgres:postgres@localhost:3002/encounter-service?sslmode=disable"
	fakeEncountersNumber = 40000
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	dbConnString := os.Getenv("POSTGRES_URI")
	if dbConnString == "" {
		dbConnString = defaultDbConnString
	}

	dbClient, err := NewPostgresClient(dbConnString)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer dbClient.Close()

	encounterPostgresStore := storage.NewEncounterPostgresStore(dbClient)
	fakeDataFlag := os.Getenv("POPULATE_FAKE_DATA")
	if fakeDataFlag == "true" {
		if err := faker.PopulateFakeEncounters(fakeEncountersNumber, encounterPostgresStore); err != nil {
			slog.Error(err.Error())
			return
		}
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{EncounterStore: encounterPostgresStore},
		}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func NewPostgresClient(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
