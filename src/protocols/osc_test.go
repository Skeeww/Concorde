package protocols_test

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/Skeeww/Concorde/src/protocols"
)

func TestWithInt32(t *testing.T) {
	msg := protocols.NewMessage("/test/a/b/c")

	msg.WithInt32(49)

	var val int32
	if err := binary.Read(bytes.NewReader(msg.Arguments), binary.BigEndian, &val); err != nil {
		t.Error("error while writing in int32 val,", err)
	}

	if msg.Type != ",i" {
		t.Errorf("wrong value, expected %s got %s", ",i", msg.Type)
	}
	if val != 49 {
		t.Errorf("wrong value, expected %d got %d", 49, val)
	}
}

func BenchmarkWithInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msg := protocols.NewMessage("/test/a/b/c")

		expectedValue := rand.Int31()
		msg.WithInt32(expectedValue)

		var val int32
		if err := binary.Read(bytes.NewReader(msg.Arguments), binary.BigEndian, &val); err != nil {
			b.Error("error while writing in int32 val,", err)
		}

		if msg.Type != ",i" {
			b.Errorf("wrong value, expected %s got %s", ",i", msg.Type)
		}
		if val != expectedValue {
			b.Errorf("wrong value, expected %d got %d", expectedValue, val)
		}
	}
}
