package protocols

import (
	"bytes"
	"encoding"
	"encoding/binary"
)

const (
	OscInt32   = "i"
	OscFloat32 = "f"
	OscString  = "s"
	OscBlob    = "b"
)

type OSCPacket struct {
	data []byte
	size int32
}

type OSCPacketer interface {
	encoding.BinaryMarshaler
}

type OSCMessage struct {
	Address   string
	Type      string
	Arguments []byte
}

func NewMessage(address string) *OSCMessage {
	return &OSCMessage{
		Address:   address,
		Type:      ",",
		Arguments: make([]byte, 0),
	}
}

func (message *OSCMessage) WithInt32(val int32) {
	buffer := bytes.Buffer{}
	if err := binary.Write(&buffer, binary.BigEndian, val); err != nil {
		return
	}
	message.Type += OscInt32
	message.Arguments = append(message.Arguments, buffer.Bytes()...)
}
