package log

func InitDBMysql(dsn string) {
	initErrorDBMysql(dsn + ".error")
	initWarnDBMysql(dsn + ".warn")
	initInfoDBMysql(dsn + ".info")
	initDebugDBMysql(dsn + ".debug")
}

func InitDBPostgres(dsn string) {
	initErrorDBPostgres(dsn + ".error")
	initWarnDBPostgres(dsn + ".warn")
	initInfoDBPostgres(dsn + ".info")
	initDebugDBPostgres(dsn + ".debug")
}

func InitDBSqlite(path string) {
	initErrorDBSqlite(path + "/error.db")
	initWarnDBSqlite(path + "/warn.db")
	initInfoDBSqlite(path + "/info.db")
	initDebugDBSqlite(path + "/debug.db")
}
