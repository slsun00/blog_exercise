package utils

import (
	"database/sql"
	"fmt"
)

var (
	Db *sql.DB
	err error
)

func init() {
	Db,err = sql.Open("mysql","root:123@tcp(localhost:3306)/bookstore0612")
	if err != nil {
		fmt.Println(" sql.Open err=",err)
	}
}
