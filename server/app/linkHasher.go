package app

import (
	"hash/fnv"
	"math/rand"
	"time"
)

type LinkHasher struct {
	letterRunes []rune
}

func (l LinkHasher) Hash(url string, numOfSymbols int) (hash string) {
	h := fnv.New64a()
	h.Write([]byte(url))
	rand.Seed(int64(h.Sum64())-time.Now().UnixNano())

	b := make([]rune, numOfSymbols)
	for i := range b {
		b[i] = l.letterRunes[rand.Intn(len(l.letterRunes))]
	}
	return string(b)
}

func NewLinkHasher() *LinkHasher {
	return &LinkHasher {
		[]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"),
	}
}
