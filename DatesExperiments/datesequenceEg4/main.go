package main

import (
	"fmt"
	"time"
)

func main() {
	// В високосном году 2016 было 366 дней.

	dateSeqArr := dateSeq("2022-03-01", "2022-03-14")
	//sort.Ints(dateSeqArr)
	fmt.Println(dateSeqArr)

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

	fmt.Println(sequence)
	//создаем массив структур

	dateSeqArr := make(map[int]dateFromdateTo)

	if 3 > len(sequence) {
		dateSeqArr[0] = dateFromdateTo{dateFrom: dateFrom, dateTo: dateTo}
	} else {

		lenSeqD := len(sequence) / 3

		fmt.Println("Длина массива")
		fmt.Println(lenSeqD)
		lenSeqR := len(sequence) % 3
		fmt.Println(lenSeqR)

		lenSeqTotal := lenSeqD + lenSeqR
		fmt.Println("Длина массива total")
		fmt.Println(lenSeqD)

		var first int = 0
		var second int = 1

		for i := 0; i < lenSeqTotal; i++ {

			//fmt.Println("Цикл")
			//fmt.Println(i)

			if i == 0 {
				//first++
				second++

				//first = 0

				//fmt.Println("Первое условие")
				//fmt.Println(first)
				//fmt.Println(second)
			} else {
				//first = second + 1
				//first += 2
				second += 3
				first = second - 2
				//first = 0

				//fmt.Println("Второе условие")
				//fmt.Println("First")
				//fmt.Println(first)
				//fmt.Println("Second")
				//fmt.Println(second)
			}

			if second < lenSeqD*3 {

				//fmt.Println("Первый вариант")
				//fmt.Println(sequence[first])
				//fmt.Println(sequence[second])

				//fmt.Println(sequence[second])
				//dateTo: sequence[second]
				dateSeqArr[i] = dateFromdateTo{dateFrom: sequence[first], dateTo: sequence[second]}

			} else {

				if lenSeqR == 1 {
					fmt.Println("Второй вариант один ")
					fmt.Println(first)
					dateSeqArr[i] = dateFromdateTo{dateFrom: sequence[first], dateTo: sequence[first]}
				}

				if lenSeqR == 2 {
					fmt.Println("Второй остаток два")

					if len(sequence) > first {
						fmt.Println(first)
						dateSeqArr[i] = dateFromdateTo{dateFrom: sequence[first], dateTo: sequence[first]}
					} else {
						fmt.Println(first - 2)
						dateSeqArr[i] = dateFromdateTo{dateFrom: sequence[first-2], dateTo: sequence[first-2]}
					}

				}

			}

		} //end for
	}

	return dateSeqArr
}
