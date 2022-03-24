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

type Db struct {
	connect *sql.DB
}

const dbType = "sqlserver"

// InitDB Open DB connection
func InitDB(finamDsn string) (*Db, error) {
	connection, err := sql.Open(dbType, finamDsn)

	dbDriver := &Db{
		connect: connection,
	}
	return dbDriver, err
}

// CloseDB Close DB connection
func (dbPtr *Db) CloseDB() error {
	return dbPtr.connect.Close()
}

// GetFinamUserDataMssql Get finam user data from ms sql DB
// add db credentials to the yaml config file something like:  finamdsn: "sqlserver://sa:password@127.0.0.1:1433?database=master&connection+timeout=30"
func (dbPtr *Db) GetFinamUserData(login string, password string) (*FinamUserData, error) {
	var userData FinamUserData

	rows, err := dbPtr.connect.Query("exec getuser @username, @password;", sql.Named("username", login), sql.Named("password", password))
	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	vals := make([]interface{}, len(columns))
	for i, _ := range columns {
		vals[i] = new(sql.RawBytes)
	}

	if rows.Next() {
		err := rows.Scan(vals...)
		if err != nil {
			return &userData, err
		}
		for idx, colName := range columns {
			switch colName {
			case "FirstName":
				userData.Fname = string(*(vals[idx].(*sql.RawBytes)))
			case "LastName":
				userData.Lname = string(*(vals[idx].(*sql.RawBytes)))
			case "MobilePhone":
				userData.Phone = string(*(vals[idx].(*sql.RawBytes)))
			case "Email":
				userData.Email = string(*(vals[idx].(*sql.RawBytes)))
			}
		}
	}

	return &userData, err
}
