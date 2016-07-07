# Parallel

[![Build Status](https://travis-ci.org/Wayt/parallel.svg)](https://travis-ci.org/Wayt/parallel) [![Go Report Card](https://goreportcard.com/badge/github.com/Wayt/parallel)](https://goreportcard.com/report/github.com/Wayt/parallel) [![GoDoc](https://godoc.org/github.com/Wayt/parallel?status.svg)](https://godoc.org/github.com/Wayt/parallel)

Parallel's Group make goroutine run & wait easy.

It automatically check if the last parameter of your function is an error, and return it through Wait()

## Installation

Install Parallel using `go get`:

    * Latest version: go get github.com/wayt/parallel

## Example

```
package main

import (
        "fmt"
        "github.com/wayt/parallel"
)

func main() {

        g := &parallel.Group{}

        g.Go(func(text string) {
                fmt.Println(text)
        }, "Hello World !")

        g.Go(func() error {
                return fmt.Errorf("something bad happened =/")
        })

        if err := g.Wait(); err != nil {
                fmt.Printf("Got an error: %v\n", err)
        }
}
```

## Contributing

Please feel free to submit issues, fork the repository and send pull requests!
