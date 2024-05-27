/*
Необходимо код примера 1 изменить так, чтобы tcp-сервер
обрабатывал подключения параллельно
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	defer listener.Close()
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			//continue
			return
		}
		go handleConn(conn)
	}
	log.Println("Завершение работы")

}

func handleConn(c net.Conn) {
	defer c.Close()
	var i int
	for {
		_, err := io.WriteString(c, fmt.Sprintf("%d\n", i))
		if err != nil {
			log.Println(err)
			return
		}
		i++
		time.Sleep(time.Second)

	}
}
