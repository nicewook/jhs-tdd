package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %v, wanted %v, given %v", got, want, numbers)
	}
}

func TestSum2(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [...]int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %v, want %v, given numbers %v", got, want, numbers)
		}
	})

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := SumSlice(numbers)
		want := 55

		if got != want {
			t.Errorf("got %v, want %v, given numbers %v", got, want, numbers)
		}
	})
}
