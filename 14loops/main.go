package main

import "fmt"

func main() {
	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("day %v \n", day)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 5 {
			goto code
		}

		fmt.Println(rougeValue)
		rougeValue++
	}

	code:
		fmt.Println("jumping to this line")
}
