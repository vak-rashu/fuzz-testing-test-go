package main

import (
	"fmt"
)

func equalByte(a, b []byte) (error, bool) {

	if len(a) != len(b) {
		return fmt.Errorf("%v", "Lenght is not equal"), false
	}

	for i := range a {
		if a[i] != b[i] {
			return nil, false
		}
	}

	return nil, true
}
