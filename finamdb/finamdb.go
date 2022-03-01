package finamdb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type FinamUserData struct {
	Name  string
	Email string
	Phone string
}

type dbCredentials struct {
	dbType   string
	host     string
	port     string
	user     string
	password string
	dbName   string
}

var dbCred dbCredentials = dbCredentials{
	dbType:   "mysql",
	host:     "10.0.3.11",
	port:     "3306",
	user:     "webuser",
	password: "webuserpass",
	dbName:   "finam_auth",
}

const ()

func GetFinamUserData(login string, password string) (FinamUserData, error) {

	var userData FinamUserData
	dsn := dbCred.user + ":" + dbCred.password + "@tcp(" + dbCred.host + ":" + dbCred.port + ")/" + dbCred.dbName
	dbConn, err := sql.Open(dbCred.dbType, dsn)
	if err != nil {
		return userData, err
	}

	err = dbConn.QueryRow("CALL getuser(?, ?)", login, password).Scan(&userData.Name, &userData.Phone, &userData.Email)

	err = dbConn.Close()
	if err != nil {
		return FinamUserData{}, err
	}

	return userData, err
}
