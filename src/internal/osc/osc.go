package osc

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"fmt"
)

const (
	OscInt32   = "i"
	OscFloat32 = "f"
	OscString  = "s"
	OscBlob    = "b"
	BufferSize = 4
)

type OSCMessage struct {
	encoding.BinaryMarshaler
	Address   string
	Type      string
	Arguments []byte
}

func NewOSCMessage(address string) *OSCMessage {
	return &OSCMessage{
		Address:   address,
		Type:      ",",
		Arguments: make([]byte, 0),
	}
}

func (message *OSCMessage) MarshalBinary() (data []byte, err error) {
	var buf []byte = []byte(message.Address)
	filling32BitsBuffer(&buf)
	buf = append(buf, []byte(message.Type)...)
	filling32BitsBuffer(&buf)
	buf = append(buf, message.Arguments...)

	if len(buf)%4 != 0 {
		return nil, fmt.Errorf("wrong binary size (should be a multiple of 4) got %d", len(buf))
	}

	return buf, nil
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
	bufferSizeTo32 := BufferSize - (bufferSize % BufferSize)

	for range bufferSizeTo32 {
		newBuffer.WriteByte(0)
	}

	*buffer = newBuffer.Bytes()
}
