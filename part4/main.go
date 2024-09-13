package main

import (
	"fmt"
)

func main() {
	var numbers [5]int

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	fmt.Println("The numbers in the array are:")
	for i, num := range numbers {
		fmt.Printf("Element %d: %d\n", i, num)
	}
}
