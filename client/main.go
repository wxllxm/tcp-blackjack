package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const address = "localhost:42069"

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			print("Error: there was an error connecting to server %s", err)
		}
		defer conn.Close()

		fmt.Printf("WXLLXM: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: there was an error getting user input %s", err.Error())
			continue
		}
		fmt.Fprintf(conn, input)
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Error: there was an error reading server response: %s", err.Error())
		}
		fmt.Printf("Server response: %s", response)
	}

}
