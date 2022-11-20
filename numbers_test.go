package main

import (
	"fmt"
	"testing"
)

func assertEqual[T comparable](t *testing.T, expected T, actual T, msg string) {
	if actual != expected {
		explanation := fmt.Sprintf("got: %v, want: %v.", actual, expected)

		var s string
		if len(msg) > 0 {
			s = msg + ", " + explanation
		} else {
			s = explanation
		}
		t.Error(s)
	}
}

func TestCreateNumbers(t *testing.T) {
	numbers := create_numbers()
	assertEqual(t, 9, len(numbers), "")
}

func TestCalculateEvens(t *testing.T) {
	numbers := create_numbers()
	evens := filter_evens(numbers)
	assertEqual(t, 4, len(evens), "length")
	assertEqual(t, 2, evens[0], "first")
	assertEqual(t, 8, evens[len(evens)-1], "last")
}

func TestMapEntries(t *testing.T) {
	m := map[int]int{3: 9}
	entries := map_entries(m)
	assertEqual(t, 1, len(entries), "length")
	assertEqual(t, 3, entries[0].key, "key")
}

func TestCalculateSum(t *testing.T) {
	sum0 := calculate_sum([]int{})
	if sum0 != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum0, 3)
	}
	sum3 := calculate_sum([]int{0, 1, 2})
	if sum3 != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum3, 3)
	}
}

func TestReverseMap(t *testing.T) {
	m2 := create_raw_map([]int{2, 3}, []int{4, 9})
	hm2 := create_hashmap(m2)
	rm2 := reverse_map(hm2)
	assertEqual(t, 3, rm2.Get(9), "Result was incorrect")
}

func TestMain(t *testing.T) {
	print_numbers()
}
