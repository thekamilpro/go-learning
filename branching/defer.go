package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("main 1")
	defer fmt.Println("defer 1")
	fmt.Println("main 2")
	defer fmt.Println("defer 2")
	// defer calls are executed in oposite order they were called

	db, _ := sql.Open("driver", "connectionString")
	defer db.Close()

	rows, _ := db.Query("bla")
	defer rows.Close()
}
