package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func changedonUserInsert(dbMysql *sql.DB, db *sql.DB) {

	start := time.Now()
	results, err1 := dbMysql.Query(`SELECT vtMTB.crmid, vtMTD.prevalue, vtMTD.postvalue
	FROM vtigercrm.vtiger_modtracker_basic vtMTB
	INNER JOIN vtigercrm.vtiger_modtracker_detail vtMTD ON (vtMTB.id = vtMTD.id)
	WHERE DATE_FORMAT(vtMTB.changedon, '%Y-%m-%d') > "2021-03-01" AND vtMTD.fieldname = 'assigned_user_id'
	ORDER BY vtMTD.id DESC`)

	if err1 != nil {
		panic(err1.Error())
		//fmt.Println(err)
	}

	var crmid int
	var prevalueP sql.NullInt16
	var postvalueP sql.NullInt16

	prevalue := cnint16(prevalueP)
	postvalue := cnint16(postvalueP)

	res, err := db.Exec(`
    drop table if exists changedonUserTable;
    create table changedonUserTable(crmid, prevalue, postvalue);
	`)

	if err != nil {
		panic(err1.Error())
		//fmt.Println(err)
	}
	fmt.Println(res)

	//insert into t values(42), (314);
	//запускаем первый запрос чтобы получить ид ответсвенных
	for results.Next() {

		err1 = results.Scan(&crmid, &prevalueP, &postvalueP)
		if err1 != nil {
			panic(err1.Error())
			//fmt.Println(err)
		}
		_, errI := db.Exec(`insert into changedonUserTable(crmid, prevalue, postvalue) values
	   (?, ?, ?)`, crmid, prevalue, postvalue)
		if errI != nil {
			panic(errI.Error())
			//fmt.Println(err)
		}

	} //end for userId

	fmt.Println(crmid)
	//fmt.Println(crmidT)
	//fmt.Println(prevalueT)
	//fmt.Println(postvalueT)
	//_, errI := db.Exec(`insert into changedonUserTable(crmid, prevalue, postvalue)
	//values(crmid, prevalue, postvalue);`)

	//_, errI := db.Exec(`insert into changedonUserTable(crmid, prevalue, postvalue) values
	//(?, ?, ?)`, crmidT, prevalueT, postvalueT)

	rows, err3 := db.Query("select count(crmid) from changedonUserTable;")
	if err3 != nil {
		//return err3
	}

	for rows.Next() {
		var crmid int
		if err3 = rows.Scan(&crmid); err3 != nil {
			//return err
		}

		fmt.Println("crmid")
		fmt.Println(crmid)
	}

	duration := time.Since(start)
	fmt.Println("Время запроса")
	fmt.Println(duration)

}

func changedonUser(db *sql.DB) string {

	var responsible string

	responsible = ""

	results, err1 := db.Query(`SELECT vtMTB.crmid, vtMTD.prevalue, vtMTD.postvalue, vtMTB.changedon
	FROM vtigercrm.vtiger_modtracker_basic vtMTB
	INNER JOIN vtigercrm.vtiger_modtracker_detail vtMTD ON (vtMTB.id = vtMTD.id)
	WHERE DATE_FORMAT(vtMTB.changedon, '%Y-%m-%d') > "2021-01-01" AND vtMTD.fieldname = 'assigned_user_id'
	ORDER BY vtMTD.id DESC`)

	if err1 != nil {
		panic(err1.Error())
		//fmt.Println(err)
	}

	var userId int32

	var prevalue sql.NullInt32
	var postvalue sql.NullInt32

	var firstName sql.NullString
	var lastName sql.NullString

	var groupname sql.NullString

	//запускаем первый запрос чтобы получить ид ответсвенных
	for results.Next() {

		err1 = results.Scan(&prevalue, &postvalue)
		if err1 != nil {
			panic(err1.Error())
			//fmt.Println(err)
		}

		if postvalue.Int32 == 70 {
			userId = prevalue.Int32
		} else {
			userId = postvalue.Int32
		}
	} //end for userId

	//проверяем на пользователя
	results2, err2 := db.Query(`SELECT first_name, last_name FROM vtigercrm.vtiger_users WHERE id = ?`, userId)

	for results2.Next() {

		err2 = results2.Scan(&firstName, &lastName)
		if err2 != nil {
			//panic(err.Error())
			fmt.Println(err2)
		}

	}

	if lastName.Valid {
		//если пользователь есть возращаем его
		if firstName.Valid {
			responsible = firstName.String + lastName.String
		} else {
			responsible = lastName.String
		}

	} else {
		// если пользователя нет проверяем на группу
		results3, err3 := db.Query(`SELECT groupname FROM vtigercrm.vtiger_groups WHERE groupid = ?`, userId)

		for results3.Next() {

			err3 = results3.Scan(&groupname)
			if err3 != nil {
				panic(err3.Error())

			}
		}

		if groupname.Valid {
			responsible = groupname.String
		}

	}

	return responsible
}
