package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const dbStr = "user=postgres password=7900 dbname=productdb sslmode=disable"

var DB, _ = sql.Open("postgres", dbStr)
