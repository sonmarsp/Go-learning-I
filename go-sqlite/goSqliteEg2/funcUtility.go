package main

import (
	"database/sql"
	"strconv"
)

/** Проверка данных из базы на нулл и присвоение переменной значения или пустой строки*/
func cnstring(s sql.NullString) string {

	var out string
	if s.Valid {
		out = s.String
	} else {
		out = " "
	}

	return out
}

/** Проверка данных из базы на нулл и присвоение переменной значения или 0*/
func cnint16(s sql.NullInt16) int16 {

	var out int16
	if s.Valid {
		out = s.Int16
	} else {
		out = 0
	}

	return out
}

/**
* Получение поля контактное лицо в зависимости от условия
 */
func contactName(firstname sql.NullString, accountname sql.NullString) string {

	contactName := ""

	if firstname.Valid {
		contactName = firstname.String
	}
	if accountname.Valid {
		contactName = accountname.String
	}

	return contactName

}

/**
* Перевод минут в день час минуты
 */
func hoursToDays(min int64) string {

	dp := min / 1440
	hp := ((min - dp*1440) / 60)
	mp := min - (dp * 1440) - (hp * 60)

	//d := strconv.Itoa(dp)
	//h := strconv.Itoa(hp)
	//m := strconv.Itoa(mp)

	d := strconv.FormatInt(dp, 10)
	h := strconv.FormatInt(hp, 10)
	m := strconv.FormatInt(mp, 10)

	out := d + " дней " + h + " часов " + m + " минут"

	return out

}

/** Перевод время обработки по нормативу 12 час в минуты*/
func priorityToMin(priority string) int64 {

	var p int64

	switch priority {

	case "12 час.":
		p = 12 * 60
	case "24 час.":
		p = 24 * 60
	case "48 час.":
		p = 48 * 60
	case "720 час.":
		p = 720 * 60
	default:
		p = 12 * 60

	}

	return p
}

/** Проверка на просроченно*/
func checkOverdue(userTime int64, priority int64) string {

	var result string

	if userTime >= priority {
		result = "Да"
	} else {
		result = "Нет"
	}

	return result

}
