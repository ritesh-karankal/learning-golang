package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Lichi"

	fmt.Println(fruitList)

	var colorList = [4]string{"blue", "green", "red", "purple"}
	fmt.Println(colorList)
}
