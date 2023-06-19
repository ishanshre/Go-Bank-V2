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
var testDB *sql.DB

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("cannot load the environment files: %s", err)
	}
	var err error
	testDB, err = sql.Open("postgres", os.Getenv("POSTGRES_CONN_STRING"))
	if err != nil {
		log.Fatalf("could not connect to database: %s", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
