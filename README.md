# Universally Unique Lexicographically Sortable Identifier
![Project status](https://img.shields.io/badge/version-0.0.3-yellow.svg)
[![Go Report Card](https://goreportcard.com/badge/imdario/go-ulid)](https://goreportcard.com/report/imdario/go-ulid)
[![GoDoc](https://godoc.org/github.com/imdario/go-ulid?status.svg)](https://godoc.org/github.com/imdario/go-ulid)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/imdario/go-ulid/master/LICENSE)

[alizain/ulid](https://github.com/alizain/ulid) port to Golang (binary format implemented).

## Why ULID?

Check out [ULID's README](https://github.com/alizain/ulid/blob/master/README.md).

## Go

I just ported it to see how fast would be the same algorithm in Go. Also, it is cryptographically secure using crypto/rand.

### Installation

```
go get github.com/imdario/go-ulid
```

### Usage

```
import (
    "github.com/imdario/go-ulid"
)

// ...

u := ulid.New()
```

### Performance

On a Intel Core 2 Duo 6600 @ 2.40 GHz, Windows 10 and Go 1.6.3:

```
BenchmarkULID-2                  1000000              1029 ns/op              16 B/op          1 allocs/op
BenchmarkEncodedULID-2           1000000              1249 ns/op              48 B/op          2 allocs/op
BenchmarkSingleEncodedULID-2    10000000               206 ns/op              32 B/op          1 allocs/op
```

Approx. 800.640 op/s, 46 times faster than [Javascript original implementation](https://github.com/alizain/ulid#performance).

#### How does it compare to UUID?

Using [google/uuid](https://github.com/google/uuid):

```
BenchmarkUUID-2                  1000000              1041 ns/op              16 B/op          1 allocs/op
BenchmarkEncodedUUID-2           1000000              1407 ns/op              64 B/op          2 allocs/op
BenchmarkSingleEncodedUUID-2     5000000               302 ns/op              48 B/op          1 allocs/op
```

go-ulid is about 12% faster than Google's UUID!
