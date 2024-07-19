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
	if bufferSize := len(msg.Arguments); bufferSize%protocols.Buffer32 != 0 {
		t.Errorf("wrong size, expected %d got %d", protocols.Buffer32-(bufferSize%protocols.Buffer32), bufferSize)
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
