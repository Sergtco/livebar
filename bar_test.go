package livebar

import (
	"testing"
	"time"
)



func TestStyle(t *testing.T) {
    arrow := NewBar(50, Arrow)
    block := NewBar(50, Block)
    arrow.Update(25)
    block.Update(25)
}

func TestCustomStyle(t *testing.T) {
    customStyle := CustomBarStyle('<', '_', '\u0000', '.', '>')
    customBar := NewBar(50, customStyle)
    customBar.Update(50)
}

func TestDynamically(t *testing.T) {
    arrow := NewBar(50, Block)
    for i := 0; i<=100; i++ {
        arrow.Update(i)
        time.Sleep(time.Millisecond * 10)
    }
}
