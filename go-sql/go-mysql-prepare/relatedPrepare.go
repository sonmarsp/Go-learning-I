package main

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

/** Инициализируем подготовленный запрос*/
func checkRelatedPrepareSql(db *sql.DB) *sql.Stmt {

	//db.SetMaxOpenConns(1)
	//db.SetMaxIdleConns(1)

	//stmt, err := db.Prepare(`SELECT COUNT(vtCF.cf_1436) AS num
	//FROM vtigercrm.vtiger_ticketcf vtCF
	//WHERE vtCF.cf_1436 = ? GROUP BY vtCF.cf_1436`)

	stmt, err := db.Prepare(`SELECT vtCF.cf_1436
	FROM vtigercrm.vtiger_ticketcf vtCF
	WHERE vtCF.cf_1436 = ?`)

	if err != nil {
		panic(err.Error())
	}

	return stmt
}

/**Считаем связанные заявки повторные обращения */
func checkRelatedPrepare(dbPrep *sql.Stmt, ticketNo string) int16 {

	split := strings.Split(ticketNo, "_")
	tn := split[1]

	//var rP sql.NullInt16
	var out int16 = 0

	res, err := dbPrep.Query(tn)

	for res.Next() {

		//err = res.Scan(&rP)

		if err != nil {
			panic(err.Error())
		}

		//out = cnint16(rP)
		out++

	}

	return out

}
