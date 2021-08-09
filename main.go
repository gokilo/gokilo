package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

var safeExit func(error)

func main() {

	restoreFunc, err := rawMode()

	safeExit = func(err error) {
		if restoreFunc != nil {
			restoreFunc()
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	defer safeExit(nil)

	if err != nil {
		safeExit(err)
	}

	r := bufio.NewReader(os.Stdin)

	for {
		ru, _, err := r.ReadRune()

		if err == io.EOF {
			break
		} else if err != nil {
			safeExit(err)
		}
		if unicode.IsControl(ru) {
			fmt.Printf("%d\r\n", ru)
		} else {
			fmt.Printf("%d (%c)\r\n", ru, ru)
		}

		if ru == 'q' {
			break
		}
	}
}
