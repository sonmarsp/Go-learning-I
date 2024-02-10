package main

import (
	"fmt"

	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	//server.SetFormat(syslog.RFC5424)
	//server.SetFormat(syslog.Automatic)
	//server.SetFormat(syslog.RFC)

	server.SetFormat(syslog.RFC3164)
	server.SetHandler(handler)
	//server.ListenUDP("localhost:5044")
	server.ListenTCP("localhost:5044")

	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			fmt.Println(logParts)
		}
	}(channel)

	server.Wait()
}
