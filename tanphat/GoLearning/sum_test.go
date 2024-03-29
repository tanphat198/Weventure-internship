package GoLearning

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){
	numbers := []int{1,2,3,4,5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given, %v", got , want, numbers)
	}
}

func TestSumAll(t *testing.T){

	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3,9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T){

	checkSums := func(t*testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got ,want)
		}
	}

	t.Run("Make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1,2,1}, []int{0,9,1})
		want := []int{3,10}

		checkSums(t, got, want)
	})

	t.Run("safety sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3,4,5})
		want := []int{0,9}

		checkSums(t, got, want)
	})
}