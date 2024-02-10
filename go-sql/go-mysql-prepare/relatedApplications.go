package main

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

/**Считаем связанные заявки повторные обращения */
func checkRelated(db *sql.DB, ticketNo string) int16 {

	//defer db.Close()
	db.SetMaxOpenConns(1)
	//На этой функции зависает

	split := strings.Split(ticketNo, "_")
	tn := split[1]

	var rP sql.NullInt16
	var out int16

	resRelated, err := db.Query(`SELECT COUNT(vtCF.cf_1436) AS num 
	FROM vtigercrm.vtiger_ticketcf vtCF 
	WHERE vtCF.cf_1436 = ? GROUP BY vtCF.cf_1436`, tn)

	for resRelated.Next() {

		err = resRelated.Scan(&rP)

		if err != nil {
			panic(err.Error())
		}
		out = cnint16(rP)

	}

	return out
}
