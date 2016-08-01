package ulid

import (
	"crypto/rand"
	"math/big"
	"time"
)

// A ULID is a 26 character Universally Unique Lexicographically Sortable Identifier
type ULID [26]byte

var (
	rander = rand.Reader // random function
	Nil    ULID
	// Crockford"s Base32
	// https://en.wikipedia.org/wiki/Base32
	encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	enclen   = int64(len(encoding))
)


func New() ULID {
	return Must(NewRandom())
}

// Must returns ulid if err is nil and panics otherwise.
func Must(ulid ULID, err error) ULID {
	if err != nil {
		panic(err)
	}
	return ulid
}

// NewRandom returns a ULID (non-binary implementation) or panics.
//
// The strength of the ULIDs is based on the strength of the crypto/rand
// package.
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

// String returns the string form of ulid, xxxxxxxxxxxxxxxxxxxxxxxxxx
// , or "" if uuid is invalid.
func (ulid ULID) String() string {
	return string(ulid[:])
}
