package controller

import "net"

type NetworkController struct{}

func (n *NetworkController) RecvMessage(c net.Conn) *NetworkMessage {
	var netMsgController NetworkMessage

	msg := netMsgController.NewNewtworkMessage()
	c.Read(msg.Buffer[0:2]) // incoming message length
	if msg.Length() == 0 {
		return nil
	}
	bytes := make([]uint8, msg.Length())
	c.Read(bytes)
	msg.Buffer = append(msg.Buffer, bytes...)

	msg.HexDump("recv")

	return msg
}

// SendMessage sends a message to the given connection.
func SendMessage(dest net.Conn, msg *NetworkMessage) {
	dest.Write(msg.Buffer)

	msg.HexDump("send")

}
