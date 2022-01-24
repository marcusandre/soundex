[![Go Report Card](https://goreportcard.com/badge/github.com/marcusandre/soundex)](https://goreportcard.com/report/github.com/marcusandre/soundex)

# go-soundex

Implementation of the [Soundex][1] phonetic algorithm in [Go][2] for indexing
names by sound, as pronounced in English.

## Install

```bash
$ go get github.com/marcusandre/go-soundex
```

## Usage

```go
package main

import (
  "fmt"

  "github.com/marcusandre/soundex"
)

func main() {
  s := soundex.New("Bacon")
  fmt.Println(s) // => B250
}
```

## Notes

We used the great [`azer/go-anglicize`][3] to get rid of most language specific
characters before creating a _Soundex_.

[1]: https://en.wikipedia.org/wiki/Soundex
[2]: https://golang.org/doc/
[3]: https://github.com/azer/go-anglicize
