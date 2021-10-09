package main

import (
	"log"
	"net"

	"openheimer.com/server_socket/utils"
)

func main() {

	port := utils.Getenv("port", ":3000")

	listener, err := net.Listen("tcp", port)
	utils.CheckError(err)
	log.Println("Server starting on port", port)

	for {
		conn, _ := listener.Accept()
		log.Println("Client connected")
		consumerClient := ConsumerClient{&conn}
		go consumerClient.Run()
	}

}
