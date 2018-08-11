# Insertion Ordered Set

[![GoDoc Reference](https://godoc.org/go.bmvs.io/orderedset?status.svg)](http://godoc.org/go.bmvs.io/orderedset)

Implementation of an insertion ordered set in Go

## Installation

```
go get go.bmvs.io/orderedset
```

## Usage

```go
package main

import (
	"go.bmvs.io/orderedset"
)

func main() {
	m := orderedset.New()
	m.Add("bar")
	// ...
}
```

See the [godoc](https://godoc.org/go.bmvs.io/orderedset) to see all the available methods with example usage.

## Development

- Run tests with `go test -race ./...`

## License

BSD-2-Clause
