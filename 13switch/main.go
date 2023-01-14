package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()	{
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("move spot 1 or you can open")
	case 2:
		fmt.Println("move spot 2")
	case 3:
		fmt.Println("move spot 3")
		fallthrough
	case 4:
		fmt.Println("move spot 4")
		fallthrough
	case 5:
		fmt.Println("move spot 5")
	case 6:
		fmt.Println("move spot 6 or you can open")
	}
}