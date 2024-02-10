package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type row struct {
	Col1       uint64
	Col4       time.Time
	Col2       string
	Col3       []uint8
	ColIgnored string
}

func main() {
	conn, err := connect()
	if err != nil {
		panic((err))
	}

	//ctx := context.Background()

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

func exec(conn driver.Conn) error {

	ctx := context.Background()

	//defer func() {
	//	conn.Exec(ctx, "DROP TABLE example")
	//}()

	if err := conn.Exec(ctx, `DROP TABLE IF EXISTS example`); err != nil {
		return err
	}
	if err := conn.Exec(ctx, `
		CREATE TABLE example (
			  Col1 UInt64
			, Col2 String
			, Col3 Array(UInt8)
			, Col4 DateTime
		) Engine = Memory
		`); err != nil {
		return err
	}

	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO example")
	if err != nil {
		return err
	}
	for i := 0; i < 1_000; i++ {
		err := batch.AppendStruct(&row{
			Col1:       uint64(i),
			Col2:       "Golang SQL database driver",
			Col3:       []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Col4:       time.Now(),
			ColIgnored: "this will be ignored",
		})
		if err != nil {
			return err
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
