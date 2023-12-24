package livebar

import (
	"fmt"
	"strings"
    "github.com/Sergtco/livebar/pkg/term"
)

type Bar struct {
	written bool
	size    int
	style   BarStyle
}

// Updates value of bar in percents format.
//
// Returns error if bar value is not in [0, 100]
func (b *Bar) Update(value int) error {
	if value < 0 || value > 100 {
		return fmt.Errorf("Invalid bar value: %d", value)
	}
	if b.written {
		fmt.Print("\x1b[1F")
		fmt.Print("\x1b[2K")
	}
	b.written = true
	termWidth, _ := term.GetWidth()
	b.size = minInt(b.size, int(termWidth))
    fmt.Println(termWidth)
	completeNum := int(float32(value) / 100. * float32(b.size))
	res := strings.Repeat(string(b.style.complete), completeNum)
	spacings := strings.Repeat(string(b.style.fill), b.size-completeNum)
	fmt.Printf("%c%s%c%s%c\n", b.style.open, res, b.style.sep, spacings, b.style.close)
	return nil
}

func NewBar(size int, style BarStyle) Bar {
	return Bar{false, size, style}
}
