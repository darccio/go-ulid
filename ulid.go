package ulid

import (
	"crypto/rand"
	"math/big"
	"time"
)

type ULID [26]byte

var (
	rander = rand.Reader
	Nil    ULID
	// Crockford"s Base32
	// https://en.wikipedia.org/wiki/Base32
	encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	enclen   = int64(len(encoding))
)

func New() ULID {
	return Must(NewRandom())
}

func Must(ulid ULID, err error) ULID {
	if err != nil {
		panic(err)
	}
	return ulid
}

func NewRandom() (ULID, error) {
	var (
		ulid ULID
	)
	err := encodeRandom(&ulid, 10)
	if err != nil {
		return Nil, err
	}
	now := time.Now().UnixNano() / int64(time.Millisecond)
	encodeTime(&ulid, now, 10)
	return ulid, err
}

func encodeTime(ulid *ULID, now int64, ln int) {
	for x := ln - 1; x >= 0; x-- {
		mod := now % enclen
		ulid[x] = encoding[mod]
		now = (now - mod) / enclen
	}
}

func encodeRandom(ulid *ULID, pivot int) (err error) {
	elbi := big.NewInt(enclen)
	for x := pivot; x < len(ulid); x++ {
		pos, err := rand.Int(rander, elbi)
		if err != nil {
			return err
		}
		ulid[x] = encoding[pos.Int64()]
	}
	return
}

func (ulid ULID) String() string {
	return string(ulid[:])
}
