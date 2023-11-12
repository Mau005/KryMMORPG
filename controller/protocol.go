package controller

import "net"

const (
	PlayerMessageTypeInfo   uint8 = 0x15
	PlayerMessageTypeCancel uint8 = 0x17
)

type ProtocolController struct {
}

func (pc *ProtocolController) SendInvalidClientVersion(c net.Conn) {
	var msgController NetworkMessage
	msg := msgController.NewNewtworkMessage()
	msg.WriteUint8(0x0a)
	msg.WriteString("Only protocol 7.60 allowed!")
	SendMessage(c, msg)
}

func (pc *ProtocolController) SendCancelMessage(c net.Conn, str string) {
	var msgController NetworkMessage
	msg := msgController.NewNewtworkMessage()
	pc.AddPlayerMessage(msg, str, PlayerMessageTypeCancel)
	SendMessage(c, msg)
}

func (pc *ProtocolController) AddPlayerMessage(msg *NetworkMessage, str string, kind uint8) {
	msg.WriteUint8(0xb4)
	msg.WriteUint8(kind)
	msg.WriteString(str)
}
