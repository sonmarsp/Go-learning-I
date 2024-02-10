package main

import (
	"fmt"
	"time"
)

func main() {
	// В високосном году 2016 было 366 дней.

	dateSeqArr := dateSeq("2022-02-01", "2022-02-12")
	//sort.Ints(dateSeqArr)
	fmt.Println(dateSeqArr)

	for k, v := range dateSeqArr {
		fmt.Println("Ключ")
		fmt.Println(k)
		fmt.Println("Значение")
		fmt.Println(v.dateFrom)
		fmt.Println(v.dateTo)
	}

}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

type dateFromdateTo struct {
	dateFrom string
	dateTo   string
}

/** Получить даты между датами */
func dateSeq(dateFrom string, dateTo string) map[int]dateFromdateTo {

	layout := "2006-01-02"
	df, err := time.Parse(layout, dateFrom)
	dt, err := time.Parse(layout, dateTo)

	if err != nil {
		panic(err)
	}

	//fmt.Println(df)
	//fmt.Println(dt)
	days := dt.Sub(df).Hours() / 24

	//fmt.Println("Разница дней")
	//fmt.Println(days)

	//dff := df.Format(layout)
	var sequence []string
	//sequence = append(sequence, dff)

	for i := 0; i <= int(days); i++ {
		//прибавляем часы которые прибавить
		var m = (24 * i)
		//прибавляем 24 часа получаем type time.time
		tempP := df.Add(time.Duration(m) * time.Hour)
		//преобразуем в строку
		temp := tempP.Format(layout)
		sequence = append(sequence, temp)
	}

	//создаем массив структур
	//focusarr := make(map[int]focusdata)
	dateSeqArr := make(map[int]dateFromdateTo)

	lenSeq := len(sequence)
	for i := 0; i < lenSeq; i++ {

		first := i
		//first := i + 1
		//second := i + 2
		second := i

		if second < lenSeq {

			//fmt.Println("Первый вариант")
			//fmt.Println(sequence[first])
			//fmt.Println(sequence[second])
			dateSeqArr[i] = dateFromdateTo{dateFrom: sequence[first], dateTo: sequence[second]}
		} else {
			//fmt.Println("Второй вариант")
			//fmt.Println(sequence[first])
			//fmt.Println(sequence[first])
			dateSeqArr[i] = dateFromdateTo{dateFrom: sequence[first], dateTo: sequence[first]}
		}

	}

	return dateSeqArr
}

/** Получить даты между датами */
func dateSeqDivision(dateFrom string, dateTo string) map[int]dateFromdateTo {

	layout := "2006-01-02"
	df, err := time.Parse(layout, dateFrom)
	dt, err := time.Parse(layout, dateTo)

	if err != nil {
		panic(err)
	}

	//fmt.Println(df)
	//fmt.Println(dt)
	days := dt.Sub(df).Hours() / 24

	//fmt.Println("Разница дней")
	//fmt.Println(days)

	//dff := df.Format(layout)
	var sequence []string
	//sequence = append(sequence, dff)

	for i := 0; i <= int(days); i++ {
		//прибавляем часы которые прибавить
		var m = (24 * i)
		//прибавляем 24 часа получаем type time.time
		tempP := df.Add(time.Duration(m) * time.Hour)
		//преобразуем в строку
		temp := tempP.Format(layout)
		sequence = append(sequence, temp)
	}

	//создаем массив структур
	//focusarr := make(map[int]focusdata)
	dateSeqArr := make(map[int]dateFromdateTo)

	lenSeq := len(sequence)
	for i := 0; i < lenSeq; i++ {

		first := i
		//first := i + 1
		//second := i + 2
		second := i

		if second < lenSeq {

			//fmt.Println("Первый вариант")
			//fmt.Println(sequence[first])
			//fmt.Println(sequence[second])
			dateSeqArr[i] = dateFromdateTo{dateFrom: sequence[first], dateTo: sequence[second]}
		} else {
			//fmt.Println("Второй вариант")
			//fmt.Println(sequence[first])
			//fmt.Println(sequence[first])
			dateSeqArr[i] = dateFromdateTo{dateFrom: sequence[first], dateTo: sequence[first]}
		}

	}

	return dateSeqArr
}
