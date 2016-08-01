package ulid

import (
	"testing"
)

func TestULID(t *testing.T) {
	ulid := New()
	if len(ulid.String()) != 26 {
		t.Fatalf("should have length 26")
	}
}

func TestEncodeTime(t *testing.T) {
	var (
		ulid ULID
		time = int64(1469918176385)
	)
	encodeTime(&ulid, time, 10)
	if string(ulid[:10]) != "01ARYZ6S41" {
		t.Fatalf("expected '01ARYZ6S41', got '%s'", string(ulid[:10]))
	}
	encodeTime(&ulid, time, 12)
	if string(ulid[:12]) != "0001ARYZ6S41" {
		t.Fatalf("expected '0001ARYZ6S41', got '%s'", string(ulid[:12]))
	}
	encodeTime(&ulid, time, 8)
	if string(ulid[:8]) != "ARYZ6S41" {
		t.Fatalf("expected 'ARYZ6S41', got '%s'", string(ulid[:8]))
	}
}

func BenchmarkULID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}
