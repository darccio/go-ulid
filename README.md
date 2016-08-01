# Universally Unique Lexicographically Sortable Identifier

[alizain/ulid](https://github.com/alizain/ulid) port to Golang. Like the original, the binary format has not been implemented.

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
BenchmarkULID-2    50000             31582 ns/op
```

Approx. 31.663 op/s, almost doubling [Javascript original implementation](https://github.com/alizain/ulid#performance).
