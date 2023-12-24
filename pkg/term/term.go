package term

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type winsize struct {
	rows    uint16
	cols    uint16
}

func GetWinsize() (*winsize, error) {
	ws := &winsize{}
    out, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
    defer out.Close()
    if err != nil {
        return nil, err
    }
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(out.Fd()),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))
	if int(retCode) == -1 {
		return nil, fmt.Errorf(errno.Error())
	}
	return ws, nil
}

func GetWidth() (uint, error) {
	ws, err := GetWinsize()
	if err != nil {
		return 0, fmt.Errorf("Error making syscall: %s", err)
	}
	return uint(ws.cols), nil
}
