package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("basic addition", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

	})
}

func ExampleAdd_1() {
	sum := Add(1, 6)
	fmt.Println(sum)
	//Output: 7
}

func ExampleAdd_2() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
