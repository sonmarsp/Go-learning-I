package main

import (
	"crypto/tls"
	"database/sql"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

//sql клиент требует tls

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
		Addr: []string{"localhost:19000"},
		Auth: clickhouse.Auth{
			Database: "logger",
			Username: "default",
			Password: "",
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout:          time.Second * 30,
		Debug:                true,
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

// Тоже tls требует
//tls: first record does not look like a TLS handshake
//exit status 1
