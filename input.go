package main

import (
	"fmt"
	"unicode"
)

func processKey(key rune) {

	if unicode.IsControl(key) {
		fmt.Printf("%d\r\n", key)
	} else {
		fmt.Printf("%d (%c)\r\n", key, key)
	}

	if key == 'q' {
		safeExit(nil)
	}
}
