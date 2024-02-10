package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:password@tcp(0.0.0.1:3306)/dbname")

	if err != nil {
		db.Close()
		panic(err.Error())
	}

	//start := time.Now()
	//r := checkRelated(db, "ОБР_544389")
	//r := checkRelated(db, "ОБР_544390")

	//fmt.Println(r)

	//duration := time.Since(start)

	//fmt.Println(duration)

	checkRelatedPrepareSql := checkRelatedPrepareSql(db)

	start2 := time.Now()
	//p := checkRelatedPrepare(checkRelatedPrepareSql, "ОБР_544389")
	p := checkRelatedPrepare(checkRelatedPrepareSql, "ОБР_544390")

	//r := checkRelated(db, "ОБР_544390")

	fmt.Println(p)

	duration2 := time.Since(start2)
	fmt.Println(duration2)

}
