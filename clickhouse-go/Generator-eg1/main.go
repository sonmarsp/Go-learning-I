package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/google/uuid"
)

//sql клиент есть несколько видов подключения

func main() {
	conn := connect()

	start := time.Now()
	exec(conn)
	duration := time.Since(start)
	fmt.Println("Время выполнения")
	fmt.Println(duration)
}

// Возвращаем форматированную дату
func GenerateDateTime(unixTime int64) string {

	return time.UnixMilli(unixTime).Format("2006-01-02 15:04:05")

}

func exec(db *sql.DB) {

	triger := "test triger"

	startTime, err := time.Parse("2006-01-02 15:04:05", "2023-10-05 14:21:06")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(startTime)
	//fmt.Println(startTime.UnixMilli())

	//timeString := GenerateDateTime(startTime.UnixMilli())
	//fmt.Println(timeString)

	for i := 0; i <= 10000; i++ {

		//здесь можно задать формулу
		s := int64(i)

		timeString := GenerateDateTime(startTime.UnixMilli() + s)

		//fmt.Println(timeString)

		uuid := uuid.NewString()
		//fmt.Println(uuid)

		_, err := db.Exec(`INSERT INTO logger.messages (timestamp, uuid, message) VALUES (?, ?, ?)`, timeString, uuid, triger)

		if err != nil {
			log.Fatal(err)
		}
	}

}

func connect() *sql.DB {

	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"0.0.0.0:19000"},
		Auth: clickhouse.Auth{
			Database: "logger",
			Username: "default",
			Password: "",
		},
		//TLS: &tls.Config{
		//	InsecureSkipVerify: true,
		//},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout:          time.Second * 30,
		Debug:                false,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "my-app", Version: "0.1"},
			},
		},
	})
	conn.SetMaxIdleConns(5)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(time.Hour)

	return conn
}
