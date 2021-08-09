package main

import (
	"bufio"
	"os"
	"unicode"
)

type KeyType int

const (
	RegKey = iota
	CtrlKey
)

type Key struct {
	KeyType KeyType
	Key     rune
}

const (
	KeyCtrlQ = 17
)

var keyReader = bufio.NewReader(os.Stdin)

func readKey() (Key, error) {
	ru, _, err := keyReader.ReadRune()
	if err != nil {
		return Key{}, err
	}

	switch {
	case unicode.IsControl(ru):
		return Key{CtrlKey, ru}, err
	default:
		return Key{RegKey, ru}, err

	}
}
