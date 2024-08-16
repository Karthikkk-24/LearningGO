package main

import "fmt"

func operators() {
	// Arithmetic Operator

	// addition
	a := 10
	b := 20
	c := a + b
	fmt.Println(c)

	// subtraction
	d := b - a
	fmt.Println(d)

	// multiplication
	e := a * b
	fmt.Println(e)

	// division
	f := b / a
	fmt.Println(f)

	// modulus
	g := a % b
	fmt.Println(g)

	// increment
	h := a
	h++
	fmt.Println(h)

	// decrement
	i := b
	i--
	fmt.Println(i)

	// assignment
	j := 10
	j += 5
	fmt.Println(j)

	// comparison
	k := 10
	l := 5
	m := k > l
	fmt.Println(m)

	// logical
	n := true
	o := false
	p := n && o
	fmt.Println(p)

	// bitwise
	q := 10
	r := 5
	s := q & r
	fmt.Println(s)

	// string concatenation
	t := "Hello"
	u := "World"
	v := t + " " + u
	fmt.Println(v)

}
