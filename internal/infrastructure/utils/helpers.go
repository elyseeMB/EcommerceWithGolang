package utils

import (
	"database/sql"
	"math/rand"
	"testing"
	"time"
)

func GenerateID() int {
	seed := time.Now().UnixNano()
	randomGenerator := rand.New(rand.NewSource(seed))
	return randomGenerator.Intn(9000) + 1000
}

func SetupMockDatabase(t *testing.T) *sql.DB {

	db, err := sql.Open("sqlite3", "./tmp/mockdb.db")

	if err != nil {
		panic(err)
	}

	return db
}
