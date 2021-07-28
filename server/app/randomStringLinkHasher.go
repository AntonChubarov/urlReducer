package app

import (
	"hash/fnv"
	"math/rand"
	"server/config"
	"time"
)

type RandomStringLinkHasher struct {
	letterRunes []rune
	numOfSymbols int
}

func (r *RandomStringLinkHasher) Hash(url string) (hash string) {
	h := fnv.New64a()
	h.Write([]byte(url))
	rand.Seed(int64(h.Sum64())-time.Now().UnixNano())

	b := make([]rune, r.numOfSymbols)
	for i := range b {
		b[i] = r.letterRunes[rand.Intn(len(r.letterRunes))]
	}
	return string(b)
}

func NewLinkHasher(sConfig *config.ServerConfig) *RandomStringLinkHasher {
	return &RandomStringLinkHasher{
		letterRunes:  []rune(sConfig.Hash.HashSymbols),
		numOfSymbols: sConfig.Hash.NumOfHashSymbols,
	}
}
