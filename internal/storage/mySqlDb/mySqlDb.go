package mysqldb

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDB struct {
	Db *sql.DB
}

func New() (*MysqlDB, error) {

	db, err := sql.Open("mysql", os.Getenv("DB_Username")+":"+os.Getenv("DB_Password")+"@tcp("+os.Getenv("DB_Host")+":"+os.Getenv("DB_Port")+")/"+os.Getenv("DB_Name"))
	if err != nil {
		return nil, err
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// create inital table if needed
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS `coursehub`")
	if err != nil {
		return nil, err
	}

	return &MysqlDB{
		Db: db,
	}, nil

}
