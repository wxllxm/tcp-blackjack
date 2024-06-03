package main

import (
	"bufio"
	"fmt"
	"net"
)

const address = "localhost:42069"

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Printf("Listening on: %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error: encountered an error accepting connection: %s\n", err.Error())
			continue
		}

		reader := bufio.NewReader(conn)

		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: there was an error reading data: %s\n", err.Error())
			return
		}

		fmt.Printf("Received: %s", data)
		conn.Write([]byte(string(data)))
		conn.Close()
	}

}
