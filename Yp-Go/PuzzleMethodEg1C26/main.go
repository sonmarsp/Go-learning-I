package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	startTime time.Time
	splits    []time.Time
}

func (sw *Stopwatch) Start() {
	sw.startTime = time.Now()
	sw.splits = nil
}

func (sw *Stopwatch) SaveSplit() {
	sw.splits = append(sw.splits, time.Now())
}

func (sw Stopwatch) GetResults() (retResult []time.Duration) {
	for _, splitTime := range sw.splits {
		retResult = append(retResult, splitTime.Sub(sw.startTime))
	}

	return
}

func main() {

	var str string = `Пример секундомера`
	fmt.Println(str)

	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())

}
