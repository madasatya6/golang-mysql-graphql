package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func GetConnection() (*sql.DB,error){
	db,err := sql.Open("mysql","user:pass@tcp(xxx.xxx.xxx.xxx:PORT)/namedb")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db,err
}
