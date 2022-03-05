package finamdb

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

type FinamUserData struct {
	Fname string
	Lname string
	Email string
	Phone string
}

type dbCredentials struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

// TODO: move credentials to config.
var dbCredMysql dbCredentials = dbCredentials{
	host:     "10.0.3.11",
	port:     "3306",
	user:     "webuser",
	password: "webuserpass",
	dbName:   "finam_auth",
}

var dbCredMssql dbCredentials = dbCredentials{
	host:     "127.0.0.1",
	port:     "1433",
	user:     "sa",
	password: "Juurfj48TGu",
	dbName:   "master",
}

func GetFinamUserDataMysql(login string, password string) (*FinamUserData, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbCredMysql.user, dbCredMysql.password, dbCredMysql.host, dbCredMysql.port, dbCredMysql.dbName)
	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	userData := &FinamUserData{}
	err = dbConn.
		QueryRow("CALL getuser(?, ?)", login, password).
		Scan(&userData.Fname, &userData.Lname, &userData.Phone, &userData.Email)
	return userData, err
}

func GetFinamUserDataMssql(login string, password string) (*FinamUserData, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		dbCredMysql.user, dbCredMysql.password, dbCredMysql.host, dbCredMysql.port, dbCredMysql.dbName)
	dbConn, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	userData := &FinamUserData{}
	err = dbConn.
		QueryRow("exec getuser @username, @password;", sql.Named("username", login), sql.Named("password", password)).
		Scan(&userData.Fname, &userData.Lname, &userData.Phone, &userData.Email)
	return userData, err
}
