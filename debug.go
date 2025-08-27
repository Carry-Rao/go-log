package log

import (
	"database/sql"

	database "github.com/Carry-Rao/go-db"
)

var debugDB *sql.DB

func initDebugDB(typ string, dsn string) {
	var err error
	debugDB, err = database.Open(typ, dsn)
	if err != nil {
		panic(err)
	}
	err = debugDB.Ping()
	if err != nil {
		panic(err)
	}

	_, err = debugDB.Exec(`
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
