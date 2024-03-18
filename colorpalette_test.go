package colorpalette

import (
	"testing"

	c "github.com/fatih/color"
)

func init() {
	c.NoColor = false
}

func TestGet(t *testing.T) {
	if !Exists("red") {
		t.Error("red doesn't exist")
	}
	col := Get("red")
	if col == nil {
		t.Error("no red")
	}
	if col.Sprint("apple") != "\x1b[91mapple\x1b[0m" {
		t.Error(col.Sprint("wrong red color"))
	}

	if len(PALETTE) != len(List()) {
		t.Error("wrong length")
	}

	for _, colorName := range List() {
		if col, ok := PALETTE[colorName]; !ok {
			t.Errorf("color %s doesn't exist in palette", colorName)
			if col().Sprintf("foo") != "" {
				t.Fatalf("color %s is wrong", colorName)
			}
		}
	}
}
