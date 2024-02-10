package main

import (
	"time"
)

/** Получить даты между датами деление на 1 день*/
func dateSeq1day(dateFrom string, dateTo string) []string {

	layout := "2006-01-02 15:04:05"
	df, err := time.Parse(layout, dateFrom)

	if err != nil {
		panic(err)
	}

	dt, err := time.Parse(layout, dateTo)
	if err != nil {
		panic(err)
	}

	days := dt.Sub(df).Hours() / 24

	var sequence []string

	for i := 0; i <= int(days); i++ {
		//прибавляем часы которые прибавить
		var m = (24 * i)
		//прибавляем 24 часа получаем type time.time
		tempP := df.Add(time.Duration(m) * time.Hour)
		//преобразуем в строку
		temp := tempP.Format(layout)
		sequence = append(sequence, temp)
	}

	return sequence
}
