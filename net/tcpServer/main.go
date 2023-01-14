package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		//reader := bufio.NewReader(conn)

		var buff [1024]byte
		//recvLen, err := reader.Read(buff[:])
		recvLen, err := conn.Read(buff[:])

		if err != nil {
			break
		}
		recvStr := string(buff[:recvLen])
		fmt.Printf("recv : %s from client : %s\n", recvStr, conn.RemoteAddr())
		conn.Write([]byte("got y"))

	}

}

func main() {

	listener, err := net.Listen("tcp", "localhost:20000")

	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		return
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("accept failed,err:%v\n", err)
			continue
		}

		go process(conn)
	}

}
