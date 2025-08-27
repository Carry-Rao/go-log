package log

func InitDBMysql(dsn string) {
	initErrorDB("mysql", dsn+".error")
	initWarnDB("mysql", dsn+".warn")
	initInfoDB("mysql", dsn+".info")
	initDebugDB("mysql", dsn+".debug")
}

func InitDBPostgres(dsn string) {
	initErrorDB("postgres", dsn+".error")
	initWarnDB("postgres", dsn+".warn")
	initInfoDB("postgres", dsn+".info")
	initDebugDB("postgres", dsn+".debug")
}

func InitDBSqlite(path string) {
	initErrorDB("sqlite3", path+"/error.db")
	initWarnDB("sqlite3", path+"/warn.db")
	initInfoDB("sqlite3", path+"/info.db")
	initDebugDB("sqlite3", path+"/debug.db")
}
