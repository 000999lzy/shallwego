package main

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestMapsAllFunc(t *testing.T) {
	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	m2 := map[string]int{
		"three": 3,
		"two":   100,
	}

	maps.Insert(m2, maps.All(m1))

	fmt.Print(m2)
}

func TestCollectFunc(t *testing.T) {
	s := []string{"one", "two", "three"}

	m1 := maps.Collect(slices.All(s))

	fmt.Print(m1)
}

func TestDeleteFunc(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	maps.DeleteFunc(m, func(k string, v int) bool {
		return v%2 != 0
	})

	fmt.Print(m)
}

func TestKeyAndValueFunc(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	keys := slices.Sorted(maps.Keys(m))
	fmt.Print(keys)

	values := slices.Sorted(maps.Values(m))
	fmt.Print(values)
}
