package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {

	restoreFunc, err := rawMode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\r\n", err)
		os.Exit(1)
	}
	defer restoreFunc()

	r := bufio.NewReader(os.Stdin)

	for {
		ru, _, err := r.ReadRune()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Error: reading key from Stdin: %s\r\n", err)
			os.Exit(1)
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
