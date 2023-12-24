package term

import (
	"fmt"
	"testing"
)


func TestSyscall(t *testing.T) {
    ws, err := GetWinsize();
    if err != nil{
        t.Fatal(err)
    }
    fmt.Println(ws)
}

func TestGetWidth(t *testing.T) {
    width, err := GetWidth()
    if err != nil {
        t.Fatal(err)
    }
    fmt.Println(width)
}

