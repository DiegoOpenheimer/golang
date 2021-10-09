package main

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"openheimer.com/server_socket/utils"
)

type ConsumerClient struct {
	Conn *net.Conn
}

func (c *ConsumerClient) Run() {
	scanner := bufio.NewScanner(*c.Conn)
	for scanner.Scan() {
		value := scanner.Text()
		log.Println("Value received: ", value)

		if strings.HasPrefix(value, "-1") {
			(*c.Conn).Close()
			break
		}
		if strings.HasPrefix(value, "down") {
			utils.Shutdown()
		}
		go c.write()
	}
}

func (c *ConsumerClient) write() {

	time.Sleep(500 * time.Millisecond)

	message := strconv.Itoa(rand.Intn(1000)) + "\n"

	(*c.Conn).Write([]byte(message))
}
