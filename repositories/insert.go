package repositories

import (
	"database/sql"
)

func AddUser(db *sql.DB, sqlString string, args ...any) (id int64, err error) {
	res, err := db.Exec(sqlString, args)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id, nil
}
