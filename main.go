package main

import (
	"fmt"
	"os"
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

	for {
		key, err := readKey()
		if err != nil {
			safeExit(err)
		}
		processKey(key)
	}
}
