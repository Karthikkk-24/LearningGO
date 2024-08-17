package main

import "fmt"

func basic_patterns() {
	i := 0
	j := 0

	for i < 5 {
		for j <= i {
			fmt.Println("*")
			j++
		}
		fmt.Println()
		j = 0
		i++
	}
}
