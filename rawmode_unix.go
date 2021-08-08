// +build linux

package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

// rawMode modifies terminal settings to enable raw mode in a platform specific
// way. It returns a function that can be invoked to restore previous settings.
func rawMode() (func(), error) {

	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TCGETS)
	if err != nil {
		return nil, fmt.Errorf("rawMode: error getting terminal flags: %w", err)
	}

	copy := *termios

	termios.Lflag = termios.Lflag &^ unix.ECHO

	if err := unix.IoctlSetTermios(unix.Stdin, unix.TCSETSF, termios); err != nil {
		return nil, fmt.Errorf("rawMode: error setting terminal flags: %w", err)
	}

	return func() {
		if err := unix.IoctlSetTermios(unix.Stdin, unix.TCSETSF, &copy); err != nil {
			fmt.Fprintf(os.Stderr, "rawMode: error restoring originl console settings: %s", err)
		}
	}, nil
}
