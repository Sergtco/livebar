package livebar

import (
	"fmt"
	"strings"
)

type Bar struct {
	written bool
	size    int
	style   Style
}

type Style struct {
	open     rune
	complete rune
	sep      rune
	fill     rune
	close    rune
}

func CustomStyle(open, complete, sep, fill, close rune) (s Style) {
	return Style{open, complete, sep, fill, close}
}

const empty = '\u0000'

var (
	Arrow Style = Style{'[', '-', '>', ' ', ']'}
	Block       = Style{empty, '█', '\u0000', '░', empty}
)

// Updates value of bar in percents format
func (b *Bar) Update(value int) {
	if b.written {
		fmt.Print("\x1b[1F")
		fmt.Print("\x1b[2K")
	}
	b.written = true
	completeNum := int(float32(value) / 100. * float32(b.size))
	res := strings.Repeat(string(b.style.complete), completeNum)
	spacings := strings.Repeat(string(b.style.fill), b.size-completeNum)
	fmt.Printf("%c%s%c%s%c\n", b.style.open, res, b.style.sep, spacings, b.style.close)
}

func NewBar(size int, style Style) Bar {
	return Bar{false, size, style}
}
