// +build unix, !windows

package consolesize

import (
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

// GetConsoleSize returns the current number of columns and rows in the active console window.
// The return value of this function is in the order of cols, rows.
func GetConsoleSize() (int, int) {
	var sz struct {
		rows    uint16
		cols    uint16
		xpixels uint16
		ypixels uint16
	}
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))
	cols := int(sz.cols)
	rows := int(sz.rows)
	if cols == 0 {
		st := os.Getenv("COLUMNS")
		if st != "" {
			n, _ := strconv.ParseInt(st, 10, 64)
			if n > 0 {
				cols = int(n)
			}
		}
	}
	return cols, rows
}
