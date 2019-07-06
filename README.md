# Insertion Ordered Set

[![Pipeline status](https://lab.bmvs.io/bs/orderedset.go/badges/master/pipeline.svg)](https://lab.bmvs.io/bs/ynab.go/commits/master) [![Pipeline status](https://ci.appveyor.com/api/projects/status/cxwct526940p1gkp/branch/master?svg=true)](https://ci.appveyor.com/project/brunomvsouza/orderedset-go/branch/master) [![Coverage report](https://lab.bmvs.io/bs/orderedset.go/badges/master/coverage.svg)](https://lab.bmvs.io/bs/orderedset.go/commits/master) [![Go Report Card](https://goreportcard.com/badge/lab.bmvs.io/bs/orderedset.go)](https://goreportcard.com/report/lab.bmvs.io/bs/orderedset.go) [![GoDoc Reference](https://godoc.org/go.bmvs.io/orderedset?status.svg)](http://godoc.org/go.bmvs.io/orderedset)

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

- Install dependencies with [`dep`](https://github.com/golang/dep)
- Run tests with `go test -race ./...`

## License

BSD-2-Clause
