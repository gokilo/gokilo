// +build linux

package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func rawMode() error {

	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TCGETS)
	if err != nil {
		return fmt.Errorf("could not fetch console settings: %s", err)
	}

	termios.Lflag = termios.Lflag &^ unix.ECHO

	if err := unix.IoctlSetTermios(unix.Stdin, unix.TCSETSF, termios); err != nil {
		return fmt.Errorf("could not set console settings: %s", err)
	}

	return nil
}
