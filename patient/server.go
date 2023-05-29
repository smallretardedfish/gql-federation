package main

import (
	"github.com/joho/godotenv"
	"github.com/smallretardedfish/gql-federation/patient/dataloaders"
	"github.com/smallretardedfish/gql-federation/patient/faker"
	"github.com/smallretardedfish/gql-federation/patient/storage"
	"golang.org/x/exp/slog"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/smallretardedfish/gql-federation/patient/graph"

	"database/sql"
	_ "github.com/lib/pq"
)

const (
	defaultPort         = "4004"
	defaultDbConnString = "postgresql://postgres:postgres@localhost:3004/patient-service?sslmode=disable"
	fakePatientsNumber  = 10000
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

	fakeDataFlag := os.Getenv("POPULATE_FAKE_DATA")
	if port == "" {
		port = defaultPort
	}

	patientPostgresStore := storage.NewPatientPostgresStore(dbClient)

	if fakeDataFlag == "true" {
		if err := faker.PopulatePatients(fakePatientsNumber, patientPostgresStore); err != nil {
			slog.Error(err.Error())
			return
		}
	}

	loaders := dataloaders.NewLoaders(patientPostgresStore)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{PatientStore: patientPostgresStore},
		}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloaders.Middleware(loaders, srv))

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
