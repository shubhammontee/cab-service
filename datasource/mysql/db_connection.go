package dbconnection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //we need this as we are using mysql databse
	//but tis all implements database/sql we rae not using  github.com/go-sql-driver/mysql
	//but we need else it will not recognize mysql
)

//DBConnectionInterface ...
type DBConnectionInterface interface {
	GetDatabaseConnection(string, string, string, string) *sql.DB
}

type dbConnection struct {
}

//NewDataBaseConnection ...
func NewDataBaseConnection() DBConnectionInterface {
	return &dbConnection{}
}

func (d dbConnection) GetDatabaseConnection(username, password, host, schema string) *sql.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema)
	var err error
	Client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {

		fmt.Println(err)
		panic(err)
	}
	return Client
}
