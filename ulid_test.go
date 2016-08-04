package ulid

import (
	"testing"
	"time"
)

func TestSetTime(t *testing.T) {
	var (
		ulid ULID
		ts   = time.Unix(1469918176, 385000000)
	)
	tm := ts.UnixNano() / int64(time.Millisecond)
	if tm != 1469918176385 {
		t.Fatalf("expected 1469918176385, got '%v'", tm)
	}
	setTime(&ulid, ts)
	et := ulid.String()
	if et[:10] != "01ARYZ6S41" {
		t.Fatalf("expected '01ARYZ6S41', got '%s'", et[:10])
	}
}

func TestSetTimeAfterSetRandom(t *testing.T) {
	var (
		ulid ULID
		ts   = time.Unix(1469918176, 385000000)
	)
	tm := ts.UnixNano() / int64(time.Millisecond)
	if tm != 1469918176385 {
		t.Fatalf("expected 1469918176385, got '%v'", tm)
	}
	setRandom(&ulid)
	setTime(&ulid, ts)
	et := ulid.String()
	if et[:10] != "01ARYZ6S41" {
		t.Fatalf("expected '01ARYZ6S41', got '%s'", et[:10])
	}
}

func BenchmarkULID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkEncodedULID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New().String()
	}
}

func BenchmarkSingleEncodedULID(b *testing.B) {
	u := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = u.String()
	}
}
