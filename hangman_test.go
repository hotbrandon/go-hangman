package main

import (
	"testing"
)

func TestOpenDictionary(t *testing.T) {
	t.Run("Open a non-existent file", func(t *testing.T) {
		_, err := OpenDictionary("test.txt")
		if err == nil {
			t.Errorf("expect error, got nil")
		}
	})

	t.Run("Open a file successfully", func(t *testing.T) {
		file, err := OpenDictionary("words_alpha.txt")

		if err != nil {
			t.Errorf("open an existing file but failed: %v", err)
		}
		if file == nil {
			t.Errorf("expect file to be not nil")
		}

		file.Close()

	})
}

func TestCheckValidAlphabet(t *testing.T) {
	t.Run("check valid alphabets", func(t *testing.T) {
		want := true
		got := CheckValidAlphabet('A')
		if want != got {
			t.Errorf("want %v, got %v", want, got)
		}
	})
	t.Run("check invalid alphabets", func(t *testing.T) {
		tests := []struct {
			name  string
			input rune
			want  bool
		}{
			{"uppercase letter", 'A', true},
			{"lowercase letter", 'z', true},
			{"boundary upper A", 'A', true},
			{"boundary upper Z", 'Z', true},
			{"boundary lower a", 'a', true},
			{"boundary lower z", 'z', true},
			{"digit", '1', false},
			{"special character", '!', false},
			{"space", ' ', false},
			{"non-ASCII letter", 'é', false}, // clarifies ASCII-only intent
		}

		for _, test := range tests {
			got := CheckValidAlphabet(test.input)
			if got != test.want {
				t.Errorf("CheckValidAlphabet(%q), got %v, want %v", test.input, got, test.want)
			}
		}
	})
}
