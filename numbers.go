package main

import (
	"fmt"

	. "github.com/goradd/maps"
	u "github.com/rjNemo/underscore"
)

// create go-raw map
func create_raw_map[K comparable, V any](keys []K, values []V) map[K]V {
	m := make(map[K]V, len(keys))
	for i, k := range keys {
		m[k] = values[i]
	}
	return m
}

// create a hashmap from a raw map
func create_hashmap[K comparable, V any](m map[K]V) Map[K, V] {
	hm := new(Map[K, V])
	for k, v := range m {
		hm.Set(k, v)
	}
	return *hm
}

type MapEntry[K, V any] struct {
	key   K
	value V
}

// get entries of a map
func map_entries[K comparable, V any](m map[K]V) []MapEntry[K, V] {
	entries := make([]MapEntry[K, V], 0, len(m))

	for k, v := range m {
		e := MapEntry[K, V]{key: k, value: v}
		entries = append(entries, e)
	}
	return entries
}

// return an array of numbers
func create_numbers() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
}

// return only even numbers
func filter_evens(numbers []int) []int {
	return u.Filter(numbers, func(n int) bool { return n%2 == 0 })
}

// calculate sum of an int array
func calculate_sum(numbers []int) int {
	return u.Reduce(numbers, func(n, acc int) int { return n + acc }, 0)
}

// create a map (V->K) from a map (K->V)
func reverse_map[K, V comparable](hm Map[K, V]) *Map[V, K] {
	f := func(k K, m *Map[V, K]) *Map[V, K] { m.Set(hm.Get(k), k); return m }
	return u.Reduce(hm.Keys(), f, new(Map[V, K]))
}

func print_numbers() {
	numbers := create_numbers()
	// filter even numbers from the slice
	evens := filter_evens(numbers)
	// square every number in the slice
	squares := u.Map(evens, func(n int) int { return n * n })
	m := create_raw_map(evens, squares)

	hm := create_hashmap(m)
	fmt.Println(hm)
	// why can't I use u.Map(map_entries(m), MapEntry::key)?
	keys := u.Map(map_entries(m), func(t MapEntry[int, int]) int { return t.key })
	values := u.Map(map_entries(m), func(t MapEntry[int, int]) int { return t.value })
	fmt.Println(m)
	fmt.Println(keys)
	fmt.Println(values)
	// reduce to the sum
	res := calculate_sum(numbers)
	fmt.Println(res)

	rev_hm := reverse_map(hm)
	fmt.Println(rev_hm)
}
