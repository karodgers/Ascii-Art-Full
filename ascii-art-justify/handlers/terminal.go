// handlers/terminal.go
package handlers

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func GetTerminalWidth() (int, error) {
	fd := int(os.Stdout.Fd())

	width, _, err := getTerminalSize(fd)
	if err != nil {
		return 0, fmt.Errorf("couldn't get terminal size: %v", err)
	}

	return width, nil
}

type terminalSize struct {
	rows    uint16
	columns uint16
}

func getTerminalSize(fd int) (int, int, error) {
	var size terminalSize

	// Make a system call to get the terminal size
	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&size)),
	)

	if err != 0 {
		return 0, 0, err
	}

	return int(size.columns), int(size.rows), nil
}
