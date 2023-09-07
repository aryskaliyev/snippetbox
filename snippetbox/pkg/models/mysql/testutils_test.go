package mysql

import (
	"database/sql"
	"io/ioutil"
	"testing"
)

func newTestDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("mysql", "test_web:pass@/test_snippetbox?parseTime=true&multiStatements=true")
	if err != nil {
		t.Fatal(err)
	}
}
