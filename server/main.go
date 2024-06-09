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

	fmt.Printf("Listening for connections on: %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error: encountered an error accepting connection: %s\n", err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		data, err := reader.ReadString('\n')

		fmt.Printf("Recieved: %s", data)

		if err != nil {
			if err.Error() == "EOF" {
				fmt.Printf("Client closed the connection\n")
			} else {
				fmt.Printf("Error: there was an error reading data %s\n", err.Error())
			}
			return
		}

		_, err = conn.Write([]byte(data))
		if err != nil {
			fmt.Printf("Error: there was an error sending data: %s\n", err.Error())
			return
		}
	}
}
