package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user input")
	reader := bufio.NewReader(os.Stdin)

	// comma ok, comma err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of the rating is %T", input)

}
