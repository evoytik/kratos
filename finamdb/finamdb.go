package finamdb

import (
	"database/sql"
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

const ()

func GetFinamUserDataMysql(login string, password string) (FinamUserData, error) {

	var userData FinamUserData
	dsn := dbCredMysql.user + ":" + dbCredMysql.password + "@tcp(" + dbCredMysql.host + ":" + dbCredMysql.port + ")/" + dbCredMysql.dbName
	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		return userData, err
	}

	err = dbConn.QueryRow("CALL getuser(?, ?)", login, password).Scan(&userData.Fname, &userData.Lname, &userData.Phone, &userData.Email)

	err = dbConn.Close()
	if err != nil {
		return FinamUserData{}, err
	}

	return userData, err
}

func GetFinamUserDataMssql(login string, password string) (FinamUserData, error) {
	var userData FinamUserData
	dsn := "sqlserver://" + dbCredMssql.user + ":" + dbCredMssql.password + "@" + dbCredMssql.host + ":" + dbCredMssql.port + "?database=" + dbCredMssql.dbName

	dbConn, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return userData, err
	}
	err = dbConn.QueryRow("exec getuser @inlogin='?', @inpass='?'", login, password).Scan(&userData.Fname, &userData.Lname, &userData.Phone, &userData.Email)
	err = dbConn.Close()
	if err != nil {
		return FinamUserData{}, err
	}

	return userData, err
}
