package main

import (
	"math/rand/v2"
)

const (
	charset       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetLength = len(charset)
)

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int()%charsetLength]
	}
	return string(b)
}
