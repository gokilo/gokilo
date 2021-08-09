package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

var globalState = struct {
	restoreTerminal func()
}{
	nil,
}

func safeExit(err error) {
	if globalState.restoreTerminal != nil {
		globalState.restoreTerminal()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func main() {

	restoreFunc, err := rawMode()
	if err != nil {
		safeExit(err)
	}
	globalState.restoreTerminal = restoreFunc
	defer safeExit(nil)

	r := bufio.NewReader(os.Stdin)

	for {
		ru, _, err := r.ReadRune()

		if err == io.EOF {
			safeExit(nil)
		} else if err != nil {
			safeExit(err)
		}
		if unicode.IsControl(ru) {
			fmt.Printf("%d\r\n", ru)
		} else {
			fmt.Printf("%d (%c)\r\n", ru, ru)
		}

		if ru == 'q' {
			safeExit(nil)
		}
	}
}
