package livebar

type BarStyle struct {
	open     rune
	complete rune
	sep      rune
	fill     rune
	close    rune
}

func CustomBarStyle(open, complete, sep, fill, close rune) (s BarStyle) {
	return BarStyle{open, complete, sep, fill, close}
}

const empty = '\u0000'

var (
	Arrow  BarStyle = BarStyle{'[', '-', '>', ' ', ']'}
	Block        = BarStyle{empty, '█', '\u0000', '░', empty}
	PacMan       = BarStyle{'[', ' ', '󰮯', '', ']'}
)

