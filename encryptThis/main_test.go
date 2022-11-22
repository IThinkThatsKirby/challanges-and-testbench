package main

import "testing"

func TestEncryptThis(t *testing.T) {
	testCases := []struct {
		arg      string
		expected string
	}{
		{arg: "A wise old owl lived in an oak", expected: "65 119esi 111dl 111lw 108dvei 105n 97n 111ka"},
	}
	for _, tC := range testCases {
		t.Run(tC.arg, func(t *testing.T) {
			if got := encryptThis(tC.arg); got != tC.expected {
				t.Error("wanted:", tC.expected, ". Instead you got:", got)
			}
		})
	}
}
