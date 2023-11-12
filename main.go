package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mau005/ServerMMORPG/controller"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:7171")
	if err != nil {
		log.Println("Error Init Socket")
		return
	}
	defer listener.Close()

	log.Println("Server Open")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var network controller.NetworkController
	var protocol controller.ProtocolController
	msg := network.RecvMessage(conn)
	if msg != nil {
		log.Println("Msg Nil")
		return
	}
	for {
		code := msg.GetBytes()
		fmt.Println("Resultado: ", code)
		switch code {
		case 0x01:
			if msg.ReadUint16() != 760 {
				protocol.SendInvalidClientVersion(conn)
			}
		}
		msg.Index++
	}

}

/*
func test(conn net.Conn) {
	defer conn.Close()
	cursor := 2
	for {
		buffer := make([]uint8, 2)
		conn.Read(buffer)
		lenBytes := binary.LittleEndian.Uint16(buffer[0:2])
		if lenBytes == 0 {
			return
		}

		bytes := make([]uint8, lenBytes)
		conn.Read(bytes)
		buffer = append(buffer, bytes...)
		fmt.Printf("\n[%s]\n%s", "recv", hex.Dump(buffer))
		if cursor >= (len(bytes))+1 {
			cursor = 0
		}
		//response := bytes[cursor]
		cursor++

	}
}
*/
