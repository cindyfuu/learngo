package main

import (
	"database/sql"
	_ "github.com/g0-sql-driver/mysql"
)

func init() {  
    sql.Register("mysql", &MySQLDriver{})
}