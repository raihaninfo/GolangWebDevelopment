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

	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println("error verifying connection with db.Ping")
	// 	panic(err.Error())
	// }


	// database insert query function
	insert, err := db.Query("INSERT INTO `product`(`id`, `name`, `price`, `description`) VALUES (`id`,'mobile','125','description')")

	if err != nil {
		fmt.Println("Insert query error", err)
	}

	insert.Close()

}
