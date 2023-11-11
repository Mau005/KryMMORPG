package controller

import (
	"encoding/binary"
	"fmt"
	"net"
)

const BUFFERSIZE = 2

const (
	networkMessageMaxSize = 16384
)

type NetworkMessage struct {
	MsgBuf  []byte
	MsgSize int
	ReadPos int
}

func NewNetworkMessage() *NetworkMessage {
	return &NetworkMessage{
		MsgBuf:  make([]byte, BUFFERSIZE),
		MsgSize: 0,
		ReadPos: 2,
	}
}

func (msg *NetworkMessage) Reset() {
	msg.MsgSize = 0
	msg.ReadPos = 2
}

func (msg *NetworkMessage) ReadFromSocket(conn net.Conn) bool {
	// Just read the size to avoid reading 2 messages at once
	sizeBytes := make([]byte, 2)
	_, err := conn.Read(sizeBytes)
	if err != nil {
		fmt.Println("Error reading size from socket:", err)
		return false
	}

	msg.MsgSize = int(binary.BigEndian.Uint16(sizeBytes))

	// Allocate buffer for the message based on the size
	msg.MsgBuf = make([]byte, msg.MsgSize)

	_, err = conn.Read(msg.MsgBuf)
	if err != nil {
		fmt.Println("Error reading message from socket:", err)
		return false
	}

	// We got something unexpected/incomplete
	if msg.MsgSize <= 2 { //|| binary.BigEndian.Uint16(msg.MsgBuf) != uint16(msg.MsgSize-2)
		msg.Reset()
		fmt.Println("We got something unexpected/incomplete")
		return false
	}

	// OK, reading starts after the size
	msg.ReadPos = 2

	return true
}

func (msg *NetworkMessage) WriteToSocket(conn net.Conn) bool {
	if msg.MsgSize == 0 {
		return true
	}

	sizeBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(sizeBytes, uint16(msg.MsgSize))

	_, err := conn.Write(sizeBytes)
	if err != nil {
		fmt.Println("Error writing size to socket:", err)
		return false
	}

	_, err = conn.Write(msg.MsgBuf[:msg.MsgSize])
	if err != nil {
		fmt.Println("Error writing message to socket:", err)
		return false
	}

	return true
}

func (msg *NetworkMessage) GetByte() byte {
	result := msg.MsgBuf[msg.ReadPos]
	msg.ReadPos++
	return result
}

func (msg *NetworkMessage) GetU16() uint16 {
	result := binary.BigEndian.Uint16(msg.MsgBuf[msg.ReadPos:])
	msg.ReadPos += 2
	return result
}

func (msg *NetworkMessage) GetU32() uint32 {
	result := binary.BigEndian.Uint32(msg.MsgBuf[msg.ReadPos:])
	msg.ReadPos += 4
	return result
}

func (msg *NetworkMessage) GetString() string {
	stringLen := int(msg.GetU16())
	if stringLen >= (networkMessageMaxSize - msg.ReadPos) {
		return ""
	}

	result := string(msg.MsgBuf[msg.ReadPos : msg.ReadPos+stringLen])
	msg.ReadPos += stringLen
	return result
}
