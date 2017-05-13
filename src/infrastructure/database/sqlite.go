package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stacknowledge/go-clean-todo/src/infrastructure"
)

type SqliteHandler struct {
	Connection *sql.DB
}

type SqliteRow struct {
	Rows *sql.Rows
}

func NewSqliteHandler(dbFile string) (*SqliteHandler, error) {
	connection, _ := sql.Open("sqlite3", dbFile)

	handler := new(SqliteHandler)
	handler.Connection = connection

	return handler, nil
}

func (handler *SqliteHandler) Execute(statement string) (int64, error) {
	result, error := handler.Connection.Exec(statement)

	if error != nil {
		return 0, error
	}

	lastInsertedID, _ := result.LastInsertId()

	return lastInsertedID, nil
}

func (handler *SqliteHandler) Query(statement string) infrastructure.Row {
	result, error := handler.Connection.Query(statement)

	rows := new(SqliteRow)

	if error != nil {
		return rows
	}

	rows.Rows = result
	return rows
}

func (row *SqliteRow) Scan(dest ...interface{}) {
	row.Rows.Scan(dest...)
}

func (row *SqliteRow) Next() bool {
	return row.Rows.Next()
}

func (row *SqliteRow) Close() error {
	return row.Rows.Close()
}
