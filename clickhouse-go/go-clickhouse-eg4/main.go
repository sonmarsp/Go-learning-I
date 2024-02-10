package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

//html

func main() {
	conn := connect()

	rows, err := conn.Query("SELECT name,toString(uuid) as uuid_str FROM system.tables LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var name, uuid string

		if err := rows.Scan(
			&name,
			&uuid,
		); err != nil {
			log.Fatal(err)
		}

		log.Printf("name: %s, uuid: %s",
			name, uuid)
	}
}

func connect() *sql.DB {

	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"127.0.0.1:8123"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 30 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Protocol: clickhouse.HTTP,
	})

	return conn
}
