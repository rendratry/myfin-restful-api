package app

import (
	"database/sql"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/myfin")
	if err != nil {
		panic(err)
	}

	db.Close()
}
