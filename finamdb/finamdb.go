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

var dbCredMssql dbCredentials = dbCredentials{
	host:     "127.0.0.1",
	port:     "1433",
	user:     "sa",
	password: "Juurfj48TGu",
	dbName:   "master",
}

// GetFinamUserDataMssql Get finam user data from ms sql DB
func GetFinamUserDataMssql(login string, password string) (*FinamUserData, error) {
	var userData FinamUserData

	//dsn := config.p.String(config.FinamDsn)
	dsn := "sqlserver://" + dbCredMssql.user + ":" + dbCredMssql.password + "@" + dbCredMssql.host + ":" + dbCredMssql.port + "?database=" + dbCredMssql.dbName + "&connection+timeout=30"

	dbConn, err := sql.Open("sqlserver", dsn)
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
