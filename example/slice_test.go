package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestSliceFunc(t *testing.T) {
	s := []string{"alice", "bob", "cell"}

	for k, v := range slices.All(s) {
		fmt.Printf("key is %d, value is %s\n", k, v)
	}

	for k, v := range slices.Backward(s) {
		fmt.Printf("back key is %d, value is %s\n", k, v)
	}

	n, f := slices.BinarySearch(s, "cell")
	fmt.Println("cell is index", n, f)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{"wang", 77},
		{"zhang", 54},
		{"liu", 60},
	}

	n, f = slices.BinarySearchFunc(people, Person{"zhang", 0}, func(a, b Person) int {
		return strings.Compare(a.name, b.name)
	})

	fmt.Println("person zhang is index", n, f)

	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sa := a[:4:10]
	fmt.Println(sa)

	seq := []int{0, 1, 1, 2, 3, 5, 8}
	fmt.Println(slices.Compact(seq))

	s1 := []int{0, 1, 2, 3}
	s2 := []int{4, 5, 6}
	fmt.Println(slices.Concat(s1, s2))

	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slices.Delete(letters, 1, 4))

	numbers := []int{0, 42, -10, 8}
	grow := slices.Grow(numbers, 2)
	fmt.Printf("the len is %d, the cap is %d\n", len(grow), cap(grow))
	mgrow := slices.Grow(grow, 5)
	fmt.Printf("the len is %d, the cap is %d\n", len(mgrow), cap(mgrow))

	oldest := slices.MaxFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println("the oldest one is", oldest.name)

	for p := range slices.Values(people) {
		fmt.Println("the name is", p.name)
	}
}
