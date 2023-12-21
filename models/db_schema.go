package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DbConnect struct {
	Db  *sql.DB
	Dsn string
}

// UserTable user表映射对象
type UserTable struct {
}
