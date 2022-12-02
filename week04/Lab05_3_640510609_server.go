package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

		// get name, output //ตรงนี้คือ รับชื่อมา
		client_name := bufio.NewReader(conn)
		isname, _ := client_name.ReadString('\n')
		isname = strings.Trim(isname, "\n") // ใช้  trim เราก็ยังไม่เข้าใจว่าคืออะไร ลองค้นดูนะ มันช่วยให้ ลบช่อง ว่างเเล้วเติมต่อกันมั่ง ลองเอาไปทำดูสู้ๆ

		go func() {

			for {
				// get message, output //อันนี้รับข้อความมา
				client_buffer := bufio.NewReader(conn)
				message, _ := client_buffer.ReadString('\n')
				fmt.Print(isname[0:len(isname)-1], " said : ", string(message))

				// create buffer to get keyboard input
				keyboard_buffer := bufio.NewReader(os.Stdin)
				fmt.Print("reply to ", isname[0:len(isname)-1], ":")
				text, _ := keyboard_buffer.ReadString('\n')

				// send to server
				fmt.Fprintf(conn, text+"\n")

			}
		}()
	}
}
