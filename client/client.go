package main

import (
	"bufio"
	"fmt"
	"gkd/utils"
	"log"
	"net"
	"os"
)

func main() {
	addr := ":8000"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if next := scanner.Scan(); next {
			line := scanner.Text()
			message := utils.NewMessage(uint32(len(line)), line)
			conn.Write(message.ToBytes())
			reader := bufio.NewReader(conn)
			if result, err := utils.NewMessageFromReader(reader); err != nil {
				log.Println(err)
			} else {
				fmt.Println(result.Data)
			}
		}
	}
}
