package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"

	"openheimer.com/client_socket/utils"
)

func main() {
	host := utils.Getenv("host", "localhost")
	port := utils.Getenv("port", ":3000")

	conn, err := net.Dial("tcp", host+port)
	utils.CheckError(err)

	log.Println("Connected in server", host+port)
	go handleReader(conn)
	handleWriter(conn)
}

func handleReader(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		value := scanner.Text()
		log.Println(value)
	}
}

func handleWriter(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		value := scanner.Text()

		conn.Write([]byte(value + "\n"))

		if strings.HasPrefix(value, "-1") || strings.HasPrefix(value, "down") {
			utils.Shutdown()
		}
	}
}
