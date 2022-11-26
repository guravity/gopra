package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArraySum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := ArraySum(numbers)
		want := 15

		if got != want {
			t.Errorf("expecetd %d but got %d, %v", want, got, numbers)
		}
	})
}

func ExampleArraySum() {
	sum := ArraySum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
	// Output: 15
}

func TestAllSumArr(t *testing.T) {
	got := AllSumArr([]int{1, 2}, []int{1, 2, 3})
	want := []int{3, 6}
	if !reflect.DeepEqual(got, want) { // 型安全ではない→コンパイルは通ってしまう
		t.Errorf("expecetd %v but got %v", want, got)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expecetd %v but got %v", want, got)
		}
	}

	t.Run("make the sums of tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 1, 9})
		want := []int{2, 10}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 1, 9})
		want := []int{0, 10}
		checkSums(t, got, want)
	})

}
