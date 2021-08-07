package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	if err := rawMode(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	r := bufio.NewReader(os.Stdin)

	for {
		b, err := r.ReadByte()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error reading from Stdin: %s\n", err)
			os.Exit(1)
		}

		if b == 'q' {
			break
		}
	}
}
