package mysqldb

import (
	"database/sql"
	"os"

	"github.com/AnkitBishen/courseHub/internal/stype"
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
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
			id int NOT NULL AUTO_INCREMENT,
			firstName varchar(100),
			lastName varchar(100),
			email varchar(500),
			password varchar(45),
			creationTime datetime DEFAULT CURRENT_TIMESTAMP,
			updateTime datetime,
			PRIMARY KEY (id)
			)`)

	if err != nil {
		return nil, err
	}

	return &MysqlDB{
		Db: db,
	}, nil

}

func (m MysqlDB) UserRegister(user stype.UserRegister) (int64, error) {
	stmt, err := m.Db.Prepare(`INSERT INTO users (firstName, lastName, email, password) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()

	return id, nil
}

func (m MysqlDB) UserLoginValidation(user stype.UserRegister) (bool, error) {
	stmt, err := m.Db.Prepare(`SELECT email FROM users WHERE email = ? AND password = ?`)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	var userData stype.UserRegister

	err = stmt.QueryRow(user.Email, user.Password).Scan(&userData.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	if user.Email != userData.Email {
		return false, nil
	} else {
		return true, nil
	}
}

func (m MysqlDB) UserValidateFromOauth(user stype.UserRegister) (bool, error) {
	stmt, err := m.Db.Prepare(`SELECT email FROM users WHERE email = ? AND password = ?`)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	var userData stype.UserRegister

	err = stmt.QueryRow(user.Email, user.Password).Scan(&userData.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	if user.Email != userData.Email {
		return false, nil
	} else {
		return true, nil
	}
}
