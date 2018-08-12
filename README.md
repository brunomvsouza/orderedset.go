# Insertion Ordered Set

[![GoDoc Reference](https://godoc.org/go.bmvs.io/orderedset?status.svg)](http://godoc.org/go.bmvs.io/orderedset) [![Build Status](https://travis-ci.com/brunomvsouza/orderedset.go.svg?branch=master)](https://travis-ci.com/brunomvsouza/orderedset.go) [![Build status](https://ci.appveyor.com/api/projects/status/cxwct526940p1gkp/branch/master?svg=true)](https://ci.appveyor.com/project/brunomvsouza/orderedset-go/branch/master) [![Coverage Status](https://coveralls.io/repos/github/brunomvsouza/orderedset.go/badge.svg?branch=master)](https://coveralls.io/github/brunomvsouza/orderedset.go)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fbrunomvsouza%2Forderedset.go.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fbrunomvsouza%2Forderedset.go?ref=badge_shield)


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


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fbrunomvsouza%2Forderedset.go.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fbrunomvsouza%2Forderedset.go?ref=badge_large)