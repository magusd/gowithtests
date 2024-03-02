package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("for each character", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertResult(t, got, want)
	})
	t.Run("for each char, specify repeat", func(t *testing.T) {
		got := Repeat("a", 7)
		want := "aaaaaaa"
		assertResult(t, got, want)
	})
	t.Run("repeat strings", func(t *testing.T) {
		got := Repeat("hello", 2)
		want := "hellohello"
		assertResult(t, got, want)
	})
}

func assertResult(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func ExampleRepeat_character() {
	fmt.Println(Repeat("x", 7))
	//Output: xxxxxxx
}
func ExampleRepeat_string() {
	fmt.Println(Repeat("hey ", 2))
	//Output: hey hey
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
