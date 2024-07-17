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
	Buffer32   = 32
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
	buffer, err := createBuffer(val)
	if err != nil {
		return
	}
	message.Type += OscInt32
	message.Arguments = append(message.Arguments, buffer.Bytes()...)
}

func (message *OSCMessage) WithFloat32(val float32) {
	buffer, err := createBuffer(val)
	if err != nil {
		return
	}
	message.Type += OscFloat32
	message.Arguments = append(message.Arguments, buffer.Bytes()...)
}

func (message *OSCMessage) WithString(val string) {
	buffer := bytes.Buffer{}
	_, err := buffer.WriteString(val)
	if err != nil {
		return
	}
	message.Type += OscString
	message.Arguments = append(message.Arguments, buffer.Bytes()...)
	filling32BitsBuffer(&message.Arguments)
}

func createBuffer(val any) (*bytes.Buffer, error) {
	buffer := bytes.Buffer{}
	if err := binary.Write(&buffer, binary.BigEndian, val); err != nil {
		return nil, err
	}
	return &buffer, nil
}

func filling32BitsBuffer(buffer *[]byte) {
	newBuffer := bytes.NewBuffer(*buffer)

	bufferSize := len(*buffer)
	bufferSizeTo32 := Buffer32 - (bufferSize % Buffer32)

	for range bufferSizeTo32 {
		newBuffer.WriteByte(0)
	}

	*buffer = newBuffer.Bytes()
}
