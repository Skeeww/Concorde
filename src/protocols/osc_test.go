package protocols_test

import (
	"bytes"
	"encoding/binary"
	"math"
	"math/rand"
	"testing"

	"github.com/Skeeww/Concorde/src/protocols"
	"github.com/Skeeww/Concorde/src/utils"
)

var (
	testMessageA = []byte("\x2f\x6f\x73\x63\x69\x6c\x6c\x61\x74\x6f\x72\x2f\x34\x2f\x66\x72\x65\x71\x75\x65\x6e\x63\x79\x00\x2c\x66\x00\x00\x43\xdc\x00\x00")
	testMessageB = []byte("\x2f\x66\x6f\x6f\x00\x00\x00\x00\x2c\x69\x69\x73\x66\x66\x00\x00\x00\x00\x03\xE8\xFF\xFF\xFF\xFF\x68\x65\x6c\x6c\x6f\x00\x00\x00\x3f\x9d\xF3\xB6\x40\xB5\xB2\x2d")
)

func TestWithInt32(t *testing.T) {
	msg := protocols.NewMessage("/test/a/b/c")

	expectedValue := rand.Int31()
	msg.WithInt32(expectedValue)

	var val int32
	if err := binary.Read(bytes.NewReader(msg.Arguments), binary.BigEndian, &val); err != nil {
		t.Error("error while writing in int32 val,", err)
	}

	if msg.Type != ",i" {
		t.Errorf("wrong value, expected %s got %s", ",i", msg.Type)
	}
	if val != expectedValue {
		t.Errorf("wrong value, expected %d got %d", expectedValue, val)
	}
}

func BenchmarkWithInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msg := protocols.NewMessage("/test/a/b/c/d")
		msg.WithInt32(math.MaxInt32)
	}
}

func TestWithFloat32(t *testing.T) {
	msg := protocols.NewMessage("/test/a/b/c")

	expectedValue := rand.Float32()
	msg.WithFloat32(expectedValue)

	var val float32
	if err := binary.Read(bytes.NewReader(msg.Arguments), binary.BigEndian, &val); err != nil {
		t.Error("error while writing in int32 val,", err)
	}

	if msg.Type != ",f" {
		t.Errorf("wrong value, expected %s got %s", ",i", msg.Type)
	}
	if val != expectedValue {
		t.Errorf("wrong value, expected %f got %f", expectedValue, val)
	}
}

func BenchmarkWithFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msg := protocols.NewMessage("/test/a/b/c")
		msg.WithFloat32(math.MaxFloat32)
	}
}

func TestWithString(t *testing.T) {
	msg := protocols.NewMessage("/test/a/b/c")

	nbChar := rand.Intn(20)
	expectedValue := utils.RandomString(nbChar)
	msg.WithString(expectedValue)

	val := string(msg.Arguments[0:nbChar])

	if msg.Type != ",s" {
		t.Errorf("wrong value, expected %s got %s", ",s", msg.Type)
	}
	if val != expectedValue {
		t.Errorf("wrong value, expected %s got %s", expectedValue, val)
	}
	if bufferSize := len(msg.Arguments); bufferSize%protocols.BufferSize != 0 {
		t.Errorf("wrong size, expected %d got %d", protocols.BufferSize-(bufferSize%protocols.BufferSize), bufferSize)
	}
}

func BenchmarkWithString26(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msg := protocols.NewMessage("/test/a/b/c")
		msg.WithString("abcdefghijklmnopqrstuvwxyz")
	}
}

func BenchmarkWithString52(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msg := protocols.NewMessage("/test/a/b/c")
		msg.WithString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
}

func TestWithMessageA(t *testing.T) {
	msg := protocols.NewMessage("/oscillator/4/frequency")

	msg.WithFloat32(440.0)

	if buf, err := msg.MarshalBinary(); !bytes.Equal(buf, testMessageA) {
		t.Errorf("wrong value, expected %x got %x", testMessageA, buf)
		if err != nil {
			t.Errorf(err.Error())
		}
		return
	}
}

func TestWithMessageB(t *testing.T) {
	msg := protocols.NewMessage("/foo")

	msg.WithInt32(1000)
	msg.WithInt32(-1)
	msg.WithString("hello")
	msg.WithFloat32(1.234)
	msg.WithFloat32(5.678)

	if buf, err := msg.MarshalBinary(); !bytes.Equal(buf, testMessageB) || err != nil {
		t.Errorf("wrong value, expected %x got %x", testMessageB, buf)
		if err != nil {
			t.Errorf(err.Error())
		}
		return
	}
}
