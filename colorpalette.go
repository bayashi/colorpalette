package colorpalette

import (
	c "github.com/fatih/color"
)

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
	"error": func() *c.Color { return c.New(c.FgHiBlue, c.BgRed, c.Underline) },
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

	"error",
}

func Get(name string) *c.Color {
	if color, ok := PALETTE[name]; ok {
		return color()
	}

	return PALETTE["error"]()
}

func List() []string {
	return orderedColors
}

func Exists(name string) bool {
	_, isExists := PALETTE[name]
	return isExists
}
