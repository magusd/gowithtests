package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("hello di", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Lobs")
		got := buffer.String()
		want := "Hello, Lobs!"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

	})
}
