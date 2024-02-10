package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Group struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Проверка подключения")

	db, err := sql.Open("mysql", "root:password@tcp(0.0.0.1:3306)/dbname")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully Connected to MySQL database")

	results, err := db.Query("SELECT groupname FROM vtiger_groups")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var groupname Group

		err = results.Scan(&groupname.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(groupname.Name)
	}

}
