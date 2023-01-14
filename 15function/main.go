package main

import "fmt"

func main() {
	result := adder(3, 5)
	fmt.Println(result)

	proRes, proStr := proAdder(2, 3, 4, 5)
	fmt.Println(proRes, proStr)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values... int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "Hi"
}

func greeter() {
	fmt.Println("hello")
}

