package main

import (
	"crypto/tls"
	"fmt"
)

func main() {

	conn, err := tls.Dial("tcp", "localhost:19000", nil)
	if err != nil {
		panic("Server doesn't support SSL certificate err: " + err.Error())
	}

	fmt.Println(conn)
}
