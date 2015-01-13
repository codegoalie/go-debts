package infrastructure

import (
	"database/sql"
	"go-debts/interfaces"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

type SqliteHandler struct {
	Conn *sql.DB
}

type SqliteRow struct {
	Rows *sql.Rows
}

func NewSqliteHandler(dbPath string) interfaces.DbHandler {
	conn, _ := sql.Open("sqlite3", dbPath)
	sqliteHandler := new(SqliteHandler)
	sqliteHandler.Conn = conn
	return sqliteHandler
}

func (handler *SqliteHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *SqliteHandler) Query(statement string) interfaces.Row {
	fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println("Error running query: ", err)
		return new(SqliteRow)
	}

	row := new(SqliteRow)
	row.Rows = rows
	return row
}

func (row *SqliteRow) Scan(dest ...interface{})  {
	row.Rows.Scan(dest...)
}

func (row *SqliteRow) Next() bool {
	return row.Rows.Next()
}
