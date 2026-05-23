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
