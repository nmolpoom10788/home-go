package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// connect to server
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")

	//get name    ตรงนี้คือเราา สร้างตัวเเปรมารับชื่อ
	keyboard_name := bufio.NewReader(os.Stdin)
	fmt.Print("Input your name:")
	getname, _ := keyboard_name.ReadString('\n') // ตรงนี้เราเทำว่าให้มันอ่านเเละส่งไปยัง server
	fmt.Fprintf(conn, getname+"\n")
	for {
		// create buffer to get keyboard input
		keyboard_buffer := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := keyboard_buffer.ReadString('\n')
		// send to server
		fmt.Fprintf(conn, text+"\n")
		// create buffer and recieve from conn
		server_buffer := bufio.NewReader(conn)
		message, _ := server_buffer.ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
