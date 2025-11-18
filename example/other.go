package main

import "fmt"

type person struct {
	name string
	age  int
}

type geometry interface {
	area() int
	perim() int
}

type rect struct {
	width, height int
}

type circle struct {
	radius int
}

func (r rect) area() int {
	return r.height * r.width
}

func (c circle) area() int {
	return c.radius * c.radius * 3
}

func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func (c circle) perim() int {
	return 2 * 3 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

type base struct {
	num int
}

func (b base) show() {
	fmt.Println("Base num:", b.num)
}

type container struct {
	base
	str string
}

func func1() {
	fmt.Println("This is a function from other.go")

	p := person{name: "Alice", age: 30}
	fmt.Println("Person:", p)

	r := rect{width: 10, height: 5}
	c := circle{radius: 7}

	measure(r)
	measure(c)

	const (
		a = iota
		b
	)

	fmt.Println("a:", a, "b:", b)

	co := container{
		base: base{num: 42},
		str:  "Hello Container",
	}

	co.show()
	fmt.Println("Container num:", co.num)
	fmt.Println("Container str:", co.str)
}
