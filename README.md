# carve

Convert PDFs into an array of PNGs. Takes a url of a PDF and converts its pages into individual PNGs.

Word documents are planned in the future.

[![BuildStatus](https://travis-ci.org/scottmotte/carve.png?branch=master)](https://travis-ci.org/scottmotte/carve)

## Usage

```Go
package main

import (
  "fmt"
  carve "github.com/scottmotte/carve"
)

func main() {
  pngs, err := carve.Convert("http://some-url.com/path-to-file.pdf")
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(pngs)
}
```
## Installation

```bash
go get github.com/scottmotte/carve
```

## Running Tests

```bash
go test -v
```
