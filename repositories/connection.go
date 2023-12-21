package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"userSystem/models"
	"userSystem/utils"
)

func Connection() (db *sql.DB, err error) {
	dbConn := models.DbConnect{Dsn: os.Getenv("DSN")}
	dbConn.Db, err = sql.Open("mysql", dbConn.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer func(Db *sql.DB) {
		err := Db.Close()
		if err != nil {
			panic(err)
		}
	}(dbConn.Db)
	return dbConn.Db, nil
}

func init() {
	utils.LoadEnv()
}
