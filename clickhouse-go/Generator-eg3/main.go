package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

// работает
type row struct {
	timestamp time.Time
	uuid      string
}

func main() {
	conn, err := connect()
	if err != nil {
		panic((err))
	}

	ctx := context.Background()

	start := time.Now()
	exec(conn, ctx)
	duration := time.Since(start)
	fmt.Println("Время выполнения")
	fmt.Println(duration)

}

// Возвращаем форматированную дату
func GenerateDateTime(unixTime int64) string {

	return time.UnixMilli(unixTime).Format("2006-01-02 15:04:05")

}

func exec(db driver.Conn, ctx context.Context) error {

	//triger := "test triger"

	//startTime, err := time.Parse("2006-01-02 15:04:05", "2023-10-05 14:21:06")
	//if err != nil {
	//	log.Fatal(err)
	//}
	batch, err := db.PrepareBatch(ctx, "INSERT INTO logger.messages")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i <= 10; i++ {

		//здесь можно задать формулу
		//s := int64(i)

		//timeString := GenerateDateTime(startTime.UnixMilli() + s)

		//fmt.Println(timeString)

		//uuid := uuid.NewString()
		//fmt.Println(uuid)

		err := batch.AppendStruct(&row{
			timestamp: time.Now(),
			uuid:      uuid.NewString(),
		})

		if err != nil {
			log.Fatal(err)
		}
	}
	return batch.Send()

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
	if err != nil {
		log.Fatal(err)
	}
	return conn, err
}
