package log

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var warnDB *sql.DB

func initWarnDBMysql(dsn string) {
	var err error
	warnDB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = warnDB.Ping()
	if err != nil {
		panic(err)
	}
	initWarnDB()
}

func initWarnDBSqlite(path string) {
	var err error
	warnDB, err = sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	err = warnDB.Ping()
	if err != nil {
		panic(err)
	}
	initWarnDB()
}

func initWarnDBPostgres(dsn string) {
	var err error
	warnDB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = warnDB.Ping()
	if err != nil {
		panic(err)
	}
	initWarnDB()
}

func initWarnDB() {
	_, err := warnDB.Exec(`
		CREATE TABLE IF NOT EXISTS warn (
			time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			message TEXT NOT NULL,
			function TEXT NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}
	_, err = warnDB.Exec(`
		CREATE TABLE IF NOT EXISTS warn_http (
			time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			url TEXT NOT NULL,
			method TEXT NOT NULL,
			status_code INT NOT NULL,
			header TEXT NOT NULL
		)
	`)
	if err != nil {
		panic(err)
	}
}

func Warn(message string, function string) {
	_, err := warnDB.Exec(`
		INSERT INTO warn (message, function) VALUES (?,?)
		`, message, function)
	if err != nil {
		panic(err)
	}
}

func WarnHTTP(url string, method string, status_code int, header string) {
	_, err := warnDB.Exec(`
		INSERT INTO warn_http (url, method, status_code, header) VALUES (?,?,?,?)
		`, url, method, status_code, header)
	if err != nil {
		panic(err)
	}
}
