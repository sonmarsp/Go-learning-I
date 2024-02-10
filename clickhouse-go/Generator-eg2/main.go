package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

//работает

func main() {
	conn, err := connect()
	if err != nil {
		panic((err))
	}

	start := time.Now()

	iterationBatch(conn, "2023-10-05 14:21:06")

	duration := time.Since(start)
	fmt.Println("Время выполнения")
	fmt.Println(duration)

}

func iterationBatch(db driver.Conn, startDateTime string) {

	for i := 0; i <= 2; i++ {

		batchInsert(db, startDateTime, 5)

	}

}

// Возвращаем форматированную дату юникс миллисекунды
func formatDateTimeMilli(unixTime int64) string {

	return time.UnixMilli(unixTime).Format("2006-01-02 15:04:05")

}

// Возвращаем форматированную дату юникс
func formatDateTimeUnix(unixTime int64) string {

	return time.Unix(unixTime, 0).Format("2006-01-02 15:04:05")

}

func batchInsert(db driver.Conn, startDateTime string, c int) {

	triger := "test triger"

	startTime, err := time.Parse("2006-01-02 15:04:05", startDateTime)
	if err != nil {
		log.Fatal(err)
	}
	batch, err := db.PrepareBatch(context.Background(), "INSERT INTO logger.messages (timestamp, uuid, message)")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i <= c; i++ {

		//здесь можно задать формулу
		s := int64(i) + rand.Int63n(100)

		timeString := formatDateTimeMilli(startTime.UnixMilli() + s)

		err := batch.Append(timeString, uuid.NewString(), triger)

		if err != nil {
			log.Fatal(err)
		}

	}

	batch.Send()

}

func connect() (driver.Conn, error) {

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:19000"},
		Auth: clickhouse.Auth{
			Database: "logger",
			Username: "default",
			Password: "",
		},
		DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
			//dialCount++
			var d net.Dialer
			return d.DialContext(ctx, "tcp", addr)
		},
		Debug: true,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format, v)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:          time.Second * 30,
		MaxOpenConns:         5,
		MaxIdleConns:         5,
		ConnMaxLifetime:      time.Duration(10) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      100,
		MaxCompressionBuffer: 102400,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "my-app", Version: "0.1"},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return conn, err
}
