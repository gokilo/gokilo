package main

import (
	"bufio"
	"os"
)

var keyReader = bufio.NewReader(os.Stdin)

func readKey() (rune, error) {
	ru, _, err := keyReader.ReadRune()
	return ru, err
}
