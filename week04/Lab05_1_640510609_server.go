package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

func main() {
	fmt.Println("Start server...")
	// listen on port 8000
	ln, err := net.Listen("tcp", ":8000")
	check(err, "Server is ready.")

	// run loop forever (or until ctrl-c)
	for {
		// accept connection
		conn, err := ln.Accept()
		check(err, "Accepted connection.")
		// get message, output
		client_buffer := bufio.NewReader(conn)
		message, _ := client_buffer.ReadString('\n')
		fmt.Print("Message Received:", string(message))

		// create buffer to get keyboard input
		keyboard_buffer := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := keyboard_buffer.ReadString('\n')

		// send to server
		fmt.Fprintf(conn, text+"\n")

	}
}
