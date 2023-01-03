package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5432 user=1315474 password=44uthy9a dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
