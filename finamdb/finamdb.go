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

// GetFinamUserDataMssql Get finam user data from ms sql DB
// add to the yaml config file something like:  finamdsn: "sqlserver://sa:password@127.0.0.1:1433?database=master&connection+timeout=30"
func GetFinamUserDataMssql(finamDsn string, login string, password string) (*FinamUserData, error) {
	var userData FinamUserData

	dbConn, err := sql.Open("sqlserver", finamDsn)
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	err = dbConn.QueryRow("exec getuser @username, @password;", sql.Named("username", login), sql.Named("password", password)).Scan(&userData.Fname, &userData.Lname, &userData.Phone, &userData.Email)
	if err != nil {
		return nil, err
	}
	err = dbConn.Close()

	return &userData, err
}
