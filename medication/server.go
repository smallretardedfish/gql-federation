package main

import (
	"github.com/smallretardedfish/gql-federation/medication/faker"
	"github.com/smallretardedfish/gql-federation/medication/storage"
	"golang.org/x/exp/slog"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/smallretardedfish/gql-federation/medication/graph"

	"database/sql"
	_ "github.com/lib/pq"
)

const (
	defaultPort           = "4003"
	defaultDbConnString   = "postgresql://postgres:postgres@0.0.0.0:3004/patient-service?sslmode=disable"
	fakeMedicationsNumber = 30000
)

func main() {
	_ = godotenv.Load()

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

	medicationPostgresStore := storage.NewMedicationPostgresStore(dbClient)

	fakeDataFlag := os.Getenv("POPULATE_FAKE_DATA")
	if fakeDataFlag == "true" {
		if err := faker.PopulateMedications(fakeMedicationsNumber, medicationPostgresStore); err != nil {
			slog.Error(err.Error())
			return
		}
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{MedicationStore: medicationPostgresStore},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Printf("db connection string: %s", dbConnString)
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
