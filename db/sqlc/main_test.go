package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("cannot load the environment files: %s", err)
	}
	conn, err := sql.Open("postgres", os.Getenv("POSTGRES_CONN_STRING"))
	if err != nil {
		log.Fatalf("could not connect to database: %s", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
