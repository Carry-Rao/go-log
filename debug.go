package log

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var debugDB *sql.DB

func initDebugDBMysql(dsn string) {
	var err error
	debugDB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = debugDB.Ping()
	if err != nil {
		panic(err)
	}
	initDebugDB()
}

func initDebugDBSqlite(path string) {
	var err error
	debugDB, err = sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	err = debugDB.Ping()
	if err != nil {
		panic(err)
	}
	initDebugDB()
}

func initDebugDBPostgres(dsn string) {
	var err error
	debugDB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = debugDB.Ping()
	if err != nil {
		panic(err)
	}
	initDebugDB()
}

func initDebugDB() {
	_, err := debugDB.Exec(`
		CREATE TABLE IF NOT EXISTS debug (
			time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			message TEXT NOT NULL,
			function TEXT NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}
	_, err = debugDB.Exec(`
		CREATE TABLE IF NOT EXISTS debug_http (
			time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			url TEXT NOT NULL,
			method TEXT NOT NULL,
			status_code INT NOT NULL
		)
	`)
	if err != nil {
		panic(err)
	}
}
