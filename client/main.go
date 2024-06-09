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

	conn, err := net.Dial("tcp", address)
	if err != nil {
		print("Error: there was an error connecting to server, check server is running...\n")
		return
	}

	defer conn.Close()

	for {
		fmt.Printf("WXLLXM: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: there was an error getting user input %s", err.Error())
			continue
		}

		if input == "exit\n" {
			fmt.Printf("closing connection...")
			break
		}

		_, err = fmt.Fprintf(conn, input)
		if err != nil {
			fmt.Printf("Error: there was an error sending input to server: %s\n", err.Error())
			continue
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Error: there was an error reading server response: %s", err.Error())
			continue
		}

		fmt.Printf("Server response: %s", response)
	}
}
