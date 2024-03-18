# colorpalette

<a href="https://github.com/bayashi/colorpalette/actions" title="colorpalette CI"><img src="https://github.com/bayashi/colorpalette/workflows/main/badge.svg" alt="colorpalette CI"></a>
<a href="https://goreportcard.com/report/github.com/bayashi/colorpalette" title="colorpalette report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/colorpalette" alt="colorpalette report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/colorpalette" title="Go colorpalette package reference" target="_blank"><img src="https://pkg.go.dev/badge/github.com/bayashi/colorpalette.svg" alt="Go Reference: colorpalette"></a>

`colorpalette` provides just a mapping of color name to [*color.Color](https://github.com/fatih/color).

## Usage

```go
package main

import (
    "fmt"

    c "github.com/bayashi/colorpalette"
)

func main() {
    col := c.Get("red")
    col.Sprint("apple") // "\x1b[91mapple\x1b[0m"
}
```

### See Also

https://github.com/fatih/color


## Installation

```cmd
go get github.com/bayashi/colorpalette
```

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
