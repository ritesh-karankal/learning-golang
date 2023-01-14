package main

import (
	"fmt"
	"sort"
)

func main() {
	fruitList := []string{"Apple", "Banana", "Peach"}
	fmt.Println(fruitList)
	fruitList = append(fruitList, "tomato")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 2
	highScores[1] = 5
	highScores[2] = 4
	highScores[3] = 6
	fmt.Println(highScores)

	highScores = append(highScores, 1, 3)
	sort.Ints(highScores)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))


	var courses = []string{"reactjs", "js", "swift", "docker", "k8s"}
	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
