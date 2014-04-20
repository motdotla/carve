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
  pngs, err := carve.Convert("http://some-url.com/path-to-file.pdf", "/local/path/to/output/dir")
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(pngs)
}
```

The output of pngs is a string of comma separated values of the path to the pngs. For example:

```
./tmp/01guest.pdf-pngs/1.png,./tmp/01guest.pdf-pngs/10.png,./tmp/01guest.pdf-pngs/11.png,./tmp/01guest.pdf-pngs/12.png,./tmp/01guest.pdf-pngs/13.png,./tmp/01guest.pdf-pngs/2.png,./tmp/01guest.pdf-pngs/3.png,./tmp/01guest.pdf-pngs/4.png,./tmp/01guest.pdf-pngs/5.png,./tmp/01guest.pdf-pngs/6.png,./tmp/01guest.pdf-pngs/7.png,./tmp/01guest.pdf-pngs/8.png,./tmp/01guest.pdf-pngs/9.png
```

## Installation

```bash
go get github.com/scottmotte/carve
```

## Running Tests

```bash
go test -v
```
