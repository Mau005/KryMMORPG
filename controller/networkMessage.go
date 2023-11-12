package controller

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type NetworkMessage struct {
	Buffer []uint8
	Index  uint16
}

func (m *NetworkMessage) NewNewtworkMessage() *NetworkMessage {
	var msg NetworkMessage
	msg.Buffer = make([]uint8, 2)
	msg.Index = 2
	return &msg
}

// Checks if cursor position is past last byte in buffer
func (p *NetworkMessage) overflow() bool {
	return p.Index >= (uint16)(len(p.Buffer))
}

// ReadUint8 reads a single byte from buffer and advances cursor.
func (p *NetworkMessage) ReadUint8() uint8 {
	if p.overflow() {
		return 0
	}
	v := p.Buffer[p.Index]
	p.Index++
	return v
}

// ReadUint16 reads 2 bytes from buffer and advances cursor.
func (p *NetworkMessage) ReadUint16() uint16 {
	if p.overflow() {
		return 0
	}
	v := binary.LittleEndian.Uint16(p.Buffer[p.Index : p.Index+2])
	p.Index += 2
	return v
}

// ReadUint32 reads 4 bytes from buffer and advances cursor.
func (p *NetworkMessage) ReadUint32() uint32 {
	if p.overflow() {
		return 0
	}
	v := binary.LittleEndian.Uint32(p.Buffer[p.Index : p.Index+4])
	p.Index += 4
	return v
}

func (p *NetworkMessage) GetBytes() (result byte) {
	result = p.Buffer[p.Index]
	p.Index++
	return result
}

// ReadString reads the string length followed by the string.
func (p *NetworkMessage) ReadString() string {
	if p.overflow() {
		return ""
	}
	var str string
	strlen := p.ReadUint16()
	for i := (uint16)(0); i < strlen; i++ {
		str += (string)(p.ReadUint8())
	}
	return str
}

// WriteUint8 writes the given byte to the message buffer. Increments message
// length by one and advances cursor.
func (p *NetworkMessage) WriteUint8(v uint8) {
	p.Buffer = append(p.Buffer, v)
	binary.LittleEndian.PutUint16(p.Buffer[0:2], (uint16)(len(p.Buffer)-2))
	p.Index++
}

// WriteUint16 writes 2 bytes to the buffer.
func (p *NetworkMessage) WriteUint16(v uint16) {
	bytes := make([]uint8, 2)
	binary.LittleEndian.PutUint16(bytes, v)
	p.WriteUint8(bytes[0])
	p.WriteUint8(bytes[1])
}

// WriteUint32 writes 4 bytes to the buffer.
func (p *NetworkMessage) WriteUint32(v uint32) {
	bytes := make([]uint8, 4)
	binary.LittleEndian.PutUint32(bytes, v)
	p.WriteUint8(bytes[0])
	p.WriteUint8(bytes[1])
	p.WriteUint8(bytes[2])
	p.WriteUint8(bytes[3])
}

// WriteString writes the string length followed by the actual string to the
// buffer.
func (p *NetworkMessage) WriteString(str string) {
	p.WriteUint16((uint16)(len(str)))
	for i := 0; i < len(str); i++ {
		p.WriteUint8((uint8)(str[i]))
	}
}

// Length returns the message length stored at the first two bytes in buffer
func (p *NetworkMessage) Length() uint16 {
	return binary.LittleEndian.Uint16(p.Buffer[0:2])
}

// SkipBytes advances the buffer by the given length n. If an overflow happens,
// the cursor returns to the previous state.
func (p *NetworkMessage) SkipBytes(n uint16) {
	p.Index += n
	if p.overflow() {
		p.Index -= n
	}
}

// HexDump is the same as hexdump -C in the terminal and is useful for debugging
func (p *NetworkMessage) HexDump(prefix string) {
	fmt.Printf("\n[%s]\n%s", prefix, hex.Dump(p.Buffer))
}
