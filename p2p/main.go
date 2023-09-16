package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	for {

		print("[0]listen\n[1]connect\n>")
		var scan string
		fmt.Scan(&scan)
		if scan == "0" {
			listen()
			break
		} else if scan == "1" {
			connect()
			break
		}

	}
	listen()

}
func connect() {
	print("host to connect:")
	var host string
	fmt.Scan(&host)
	print("port to connect:")
	var port string
	fmt.Scan(&port)
	ip, err := net.LookupIP(host)
	if err != nil {
		panic(err)
	}
	cleanip := ip[0].String() + ":" + port
	fmt.Println("trying to connect to " + cleanip)
	connection, err2 := net.Dial("tcp", cleanip)
	if err2 != nil {
		panic(err)
	}
	fmt.Println("connected")
	read(connection)
	write(connection)

}
func listen() {
	print("port to listen:")
	var port string
	fmt.Scan(&port)
	fmt.Println("Starting...")

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()
	defer fmt.Println("Listener off")

	fmt.Println("listening:")

	connection, err := listener.Accept()
	remoteIp := connection.RemoteAddr().String()
	fmt.Println("alive")
	fmt.Println(remoteIp + " is Connected!!")
	if err != nil {
		log.Fatal(err)
	}
	read(connection)
	write(connection)

}
func read(connection net.Conn) {
	go func(conn net.Conn) {
		for {
			buffer := make([]byte, 2048)
			_, err := io.ReadAtLeast(conn, buffer, 1)
			if err != nil {
				fmt.Println("Connection Closed")
				os.Exit(0)
			}
			fmt.Print("\n~" + string(buffer) + "\n>")
		}
	}(connection)
}
func write(connection net.Conn) {

	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print(">")
		scanner.Scan()
		message := scanner.Text()

		connection.Write([]byte(message))

	}

}
