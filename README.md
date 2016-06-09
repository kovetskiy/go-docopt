godoc
=========

[![Build Status](https://travis-ci.org/godoc/godoc.go.svg?branch=master)](https://travis-ci.org/godoc/godoc.go)
[![Coverage Status](https://coveralls.io/repos/godoc/godoc.go/badge.png)](https://coveralls.io/r/godoc/godoc.go)
[![GoDoc](https://godoc.org/github.com/godoc/godoc.go?status.png)](https://godoc.org/github.com/godoc/godoc.go)

An implementation of [godoc](http://godoc.org/) in the
[Go](http://golang.org/) programming language.

**godoc** helps you create *beautiful* command-line interfaces easily:

```go
package main

import (
	"fmt"
	"github.com/godoc/godoc"
)

func main() {
	  usage := `Naval Fate.

Usage:
  naval_fate ship new <name>...
  naval_fate ship <name> move <x> <y> [--speed=<kn>]
  naval_fate ship shoot <x> <y>
  naval_fate mine (set|remove) <x> <y> [--moored|--drifting]
  naval_fate -h | --help
  naval_fate --version

Options:
  -h --help     Show this screen.
  --version     Show version.
  --speed=<kn>  Speed in knots [default: 10].
  --moored      Moored (anchored) mine.
  --drifting    Drifting mine.`

	  arguments, _ := godoc.Parse(usage, nil, true, "Naval Fate 2.0", false)
	  fmt.Println(arguments)
}
```

**godoc** parses command-line arguments based on a help message. Don't
write parser code: a good help message already has all the necessary
information in it.

## Installation

⚠ Use the alias “godoc”. To use godoc in your Go code:

```go
import "github.com/godoc/godoc"
```

To install godoc according to your `$GOPATH`:

```console
$ go get github.com/godoc/godoc
```

## API

```go
func Parse(doc string, argv []string, help bool, version string,
    optionsFirst bool, exit ...bool) (map[string]interface{}, error)
```
Parse `argv` based on the command-line interface described in `doc`.

Given a conventional command-line help message, godoc creates a parser and
processes the arguments. See
https://github.com/godoc/godoc#help-message-format for a description of the
help message format. If `argv` is `nil`, `os.Args[1:]` is used.

godoc returns a map of option names to the values parsed from `argv`, and an
error or `nil`.

More documentation for godoc is available at
[GoDoc.org](https://godoc.org/github.com/godoc/godoc.go).

## Testing

All tests from the Python version are implemented and passing
at [Travis CI](https://travis-ci.org/godoc/godoc.go). New
language-agnostic tests have been added
to [test_golang.godoc](test_golang.godoc).

To run tests for godoc, use `go test`.
