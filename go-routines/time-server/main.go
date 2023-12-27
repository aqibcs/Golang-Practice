package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
	}

	func() {
		for {
			_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
			if err != nil {
				return
			}
			time.Sleep(time.Second)
		}
	}()

}
