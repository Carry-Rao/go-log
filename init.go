package log

func initDBMysql(dsn string) {
	initErrorDB("mysql", dsn+".error")
	initWarnDB("mysql", dsn+".warn")
	initInfoDB("mysql", dsn+".info")
	initDebugDB("mysql", dsn+".debug")
}

func initDBPostgres(dsn string) {
	initErrorDB("postgres", dsn+".error")
	initWarnDB("postgres", dsn+".warn")
	initInfoDB("postgres", dsn+".info")
	initDebugDB("postgres", dsn+".debug")
}

func initDBSqlite(path string) {
	initErrorDB("sqlite3", path+"/error.db")
	initWarnDB("sqlite3", path+"/warn.db")
	initInfoDB("sqlite3", path+"/info.db")
	initDebugDB("sqlite3", path+"/debug.db")
}

func InitDB(typ, path string) {
	switch typ {
	case "mysql":
		initDBMysql(path)
	case "postgres":
		initDBPostgres(path)
	case "sqlite3":
		initDBSqlite(path)
	default:
		panic("unsupported database type")
	}
}
