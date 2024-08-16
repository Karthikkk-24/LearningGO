package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var name string = "Karthik" // string declaration // output => Karthik
	var age int = 22            // int declaration // output => 22
	var isMarried bool = false  // bool declaration // output => false

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)

	// print the type of the variable
	fmt.Printf("%T\n", name)      // output => string
	fmt.Printf("%T\n", age)       // output => int
	fmt.Printf("%T\n", isMarried) // output => bool

	comments()

}
