package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checkSum := func(t *testing.T, got int, want int, numbers []int) {
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	}

	t.Run("sum of array of ints", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		checkSum(t, got, want, numbers)
	})

	t.Run("sum array of different size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		checkSum(t, got, want, numbers)

	})

}

func ExampleSum() {
	fmt.Println(Sum([]int{1, 2, 3}))
	//Output: 6
}

func TestSumAll(t *testing.T) {
	t.Run("sum many slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{6, 15}
		checkSums(t, got, want)
	})
}

func ExampleSumAll() {
	fmt.Println(SumAll([]int{1, 2, 3}, []int{4, 5, 6}))
	//Output: [6 15]
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum all last elements", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 5, 6, 9})
		want := []int{2, 20}
		checkSums(t, got, want)
	})
	t.Run("sum with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		checkSums(t, got, want)
	})
}

func checkSums(t *testing.T, got []int, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func ExampleSumAllTails() {
	fmt.Println(SumAllTails([]int{1, 2, 3}, []int{4, 5, 6}))
	//Output: [5 11]
}
