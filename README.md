# Universally Unique Lexicographically Sortable Identifier

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

On a Intel Core 2 Duo 6600 @ 2.40 GHz and Windows 10:

```
BenchmarkULID-2                  1000000              1045 ns/op              16 B/op          1 allocs/op
BenchmarkEncodedULID-2           1000000              1300 ns/op              48 B/op          2 allocs/op
BenchmarkSingleEncodedULID-2     5000000               236 ns/op              32 B/op          1 allocs/op
```

Approx. 538.500 op/s, 30 times faster than [Javascript original implementation](https://github.com/alizain/ulid#performance).

#### How does it compare to UUID?

Using [google/uuid](https://github.com/google/uuid)

```
BenchmarkUUID-2                  1000000              1041 ns/op              16 B/op          1 allocs/op
BenchmarkEncodedUUID-2           1000000              1407 ns/op              64 B/op          2 allocs/op
BenchmarkSingleEncodedUUID-2     5000000               302 ns/op              48 B/op          1 allocs/op
```

go-ulid is slightly faster than Google's UUID!
