package main

import (
	"database/sql"
	"github.com/smallretardedfish/gql-federation/condition/faker"
	"github.com/smallretardedfish/gql-federation/condition/storage"
	"golang.org/x/exp/slog"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/smallretardedfish/gql-federation/condition/graph"

	_ "github.com/lib/pq"
)

const (
	defaultPort          = "4001"
	defaultDbConnString  = "postgresql://postgres:postgres@localhost:3001/condition-service?sslmode=disable"
	fakeConditionsNumber = 18000
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

	conditionPostgresStore := storage.NewConditionPostgresStore(dbClient)

	fakeDataFlag := os.Getenv("POPULATE_FAKE_DATA")
	if fakeDataFlag == "true" {
		if err := faker.PopulateFakeConditions(fakeConditionsNumber, conditionPostgresStore); err != nil {
			slog.Error(err.Error())
			return
		}
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{ConditionStore: conditionPostgresStore},
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
