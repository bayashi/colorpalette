package colorpalette

import (
	"fmt"

	c "github.com/fatih/color"
)

const STR_ERROR = "error"

var PALETTE = map[string]func() *c.Color{
	// Bright colors
	"red":        func() *c.Color { return c.New(c.FgHiRed) },
	"green":      func() *c.Color { return c.New(c.FgHiGreen) },
	"yellow":     func() *c.Color { return c.New(c.FgHiYellow) },
	"blue":       func() *c.Color { return c.New(c.FgHiBlue) },
	"magenta":    func() *c.Color { return c.New(c.FgHiMagenta) },
	"cyan":       func() *c.Color { return c.New(c.FgHiCyan) },
	"gray":       func() *c.Color { return c.New(c.FgHiBlack) },
	"light_gray": func() *c.Color { return c.New(c.FgWhite) },
	"white":      func() *c.Color { return c.New(c.FgHiWhite) },

	// Dark colors
	"black":        func() *c.Color { return c.New(c.FgBlack) },
	"dark_red":     func() *c.Color { return c.New(c.FgRed) },
	"dark_green":   func() *c.Color { return c.New(c.FgGreen) },
	"dark_yellow":  func() *c.Color { return c.New(c.FgYellow) },
	"dark_blue":    func() *c.Color { return c.New(c.FgBlue) },
	"dark_magenta": func() *c.Color { return c.New(c.FgMagenta) },
	"dark_cyan":    func() *c.Color { return c.New(c.FgCyan) },

	// White text on background color
	"bg_red":     func() *c.Color { return c.New(c.BgRed) },
	"bg_green":   func() *c.Color { return c.New(c.BgGreen) },
	"bg_yellow":  func() *c.Color { return c.New(c.BgYellow) },
	"bg_blue":    func() *c.Color { return c.New(c.BgBlue) },
	"bg_magenta": func() *c.Color { return c.New(c.BgMagenta) },
	"bg_cyan":    func() *c.Color { return c.New(c.BgCyan) },

	// error color
	STR_ERROR: func() *c.Color { return c.New(c.FgHiBlue, c.BgRed, c.Underline) },
}

var PALETTE_CODE = map[string]string{
	// Bright colors
	"red":        fmt.Sprint(c.FgHiRed),
	"green":      fmt.Sprint(c.FgHiGreen),
	"yellow":     fmt.Sprint(c.FgHiYellow),
	"blue":       fmt.Sprint(c.FgHiBlue),
	"magenta":    fmt.Sprint(c.FgHiMagenta),
	"cyan":       fmt.Sprint(c.FgHiCyan),
	"gray":       fmt.Sprint(c.FgHiBlack),
	"light_gray": fmt.Sprint(c.FgWhite),
	"white":      fmt.Sprint(c.FgHiWhite),

	// Dark colors
	"black":        fmt.Sprint(c.FgBlack),
	"dark_red":     fmt.Sprint(c.FgRed),
	"dark_green":   fmt.Sprint(c.FgGreen),
	"dark_yellow":  fmt.Sprint(c.FgYellow),
	"dark_blue":    fmt.Sprint(c.FgBlue),
	"dark_magenta": fmt.Sprint(c.FgMagenta),
	"dark_cyan":    fmt.Sprint(c.FgCyan),

	// White text on background color
	"bg_red":     fmt.Sprint(c.BgRed),
	"bg_green":   fmt.Sprint(c.BgGreen),
	"bg_yellow":  fmt.Sprint(c.BgYellow),
	"bg_blue":    fmt.Sprint(c.BgBlue),
	"bg_magenta": fmt.Sprint(c.BgMagenta),
	"bg_cyan":    fmt.Sprint(c.BgCyan),

	// error color
	STR_ERROR: fmt.Sprint(c.BgHiRed),
}

var orderedColors = []string{
	"red",
	"green",
	"yellow",
	"blue",
	"magenta",
	"cyan",
	"gray",
	"light_gray",
	"white",

	"black",
	"dark_red",
	"dark_green",
	"dark_yellow",
	"dark_blue",
	"dark_magenta",
	"dark_cyan",

	"bg_red",
	"bg_green",
	"bg_yellow",
	"bg_blue",
	"bg_magenta",
	"bg_cyan",

	STR_ERROR,
}

func Get(name string) *c.Color {
	if color, ok := PALETTE[name]; ok {
		return color()
	}

	return PALETTE[STR_ERROR]()
}

func List() []string {
	return orderedColors
}

func Exists(name string) bool {
	_, isExists := PALETTE[name]
	return isExists
}

func GetCode(name string) string {
	if color, ok := PALETTE_CODE[name]; ok {
		return color
	}

	return PALETTE_CODE[STR_ERROR]
}
