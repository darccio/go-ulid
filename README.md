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
BenchmarkULID-2                  1000000              1416 ns/op
BenchmarkEncodedULID-2           1000000              1857 ns/op
BenchmarkSingleEncodedULID-2     5000000               380 ns/op
```

Approx. 538.500 op/s, 30 times faster than [Javascript original implementation](https://github.com/alizain/ulid#performance).
