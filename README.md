# numtostr

[![godoc.org][godoc-badge]][godoc]

`numtostr` find wrong usage of fmt.Sprint and arguments are only number.

```go
package main

import "fmt"

func main() {
    fmt.Sprint(1, 2, 3)
}
```

```sh
$ go vet -vettool=`which numtostr` main.go
./main.go:6:2: don't use fmt.Sprint to convert number to string. Use strconv.Itoa.
```

<!-- links -->
[godoc]: https://godoc.org/github.com/gostaticanalysis/numtostr
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=flat-square&label=%20godoc.org
