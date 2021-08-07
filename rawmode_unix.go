// +build linux

package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func rawMode() error {

	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TCGETS)
	if err != nil {
		return fmt.Errorf("rawMode: error getting terminal flags: %w", err)
	}

	termios.Lflag = termios.Lflag &^ unix.ECHO

	if err := unix.IoctlSetTermios(unix.Stdin, unix.TCSETSF, termios); err != nil {
		return fmt.Errorf("rawMode: error setting terminal flags: %w", err)
	}

	return nil
}
