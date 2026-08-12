package main

import (
	"testing"
)

func FuzzTest(f *testing.F) {
	tests := []string{
		"Raji",
		"Kuldeep",
		"KD",
		"JJ",
		"Gangnam",
	}

	for _, val := range tests {
		f.Add(val)
	}

	f.Fuzz(func(t *testing.T, a string) {
		// for _, val := range tests {
		// 	if got, want := name(a), val; got != want {
		// 		t.Errorf("Wanted %s, got %s", want, got)

		// 	}
		// }
		name(a)
	})
}

var seedCorpus = struct {
	a []byte
	b []byte
}{
	a: []byte{1, 2},
	b: []byte{1, 3, 5, 5, 8, 9, 6, 6},
}

func FuzzEqualTest(f *testing.F) {
	f.Add(seedCorpus.a, seedCorpus.b)

	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		equalByte(a, b)
	})
}
