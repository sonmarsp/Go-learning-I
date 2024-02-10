package main

import (
	"log"
	"log/syslog"
)

func main() {
	s, err := syslog.New(syslog.LOG_WARNING|syslog.LOG_LOCAL7, "test custom")
	if err != nil {
		log.Fatal(err)
	}

	s.Emerg("log test proverka")
}
