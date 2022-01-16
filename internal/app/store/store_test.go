package store_test

import (
	"os"
	"testing"
)

var databaseURL string

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=8081 dbname=restapi_test user=task password=task sslmode=disable"
	}

	os.Exit(m.Run())
}
