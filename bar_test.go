package livebar

import (
	"testing"
	"time"
)



func TestStyle(t *testing.T) {
    arrow := newBar(50, Arrow)
    block := newBar(50, Block)
    arrow.Update(25)
    block.Update(25)
}

func TestCustomStyle(t *testing.T) {
    customStyle := CustomStyle('<', '_', '\u0000', '.', '>')
    customBar := newBar(50, customStyle)
    customBar.Update(50)
}

func TestDynamically(t *testing.T) {
    arrow := newBar(50, Block)
    for i := 0; i<=100; i++ {
        arrow.Update(i)
        time.Sleep(time.Millisecond * 10)
    }
}
