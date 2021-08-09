package main

import (
	"fmt"
)

func processKey(key Key) {

	switch key.KeyType {
	case CtrlKey:
		fmt.Printf("Control Key: %d\r\n", key.Key)
	case RegKey:
		fmt.Printf("Regular Key: %d (%c)\r\n", key.Key, key.Key)
	}

	if key.KeyType == CtrlKey && key.Key == KeyCtrlQ {
		safeExit(nil)
	}
}
