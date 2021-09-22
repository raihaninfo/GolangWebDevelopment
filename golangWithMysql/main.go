package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/sql-golang")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database....")
	defer db.Close()
}
