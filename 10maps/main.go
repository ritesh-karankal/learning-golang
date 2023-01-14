package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"


	fmt.Println("The list of languages are", languages)
	fmt.Println("JS short for", languages["JS"])

	delete(languages, "RB")
	fmt.Println("The list of languages are", languages)


	for _, value := range languages{
		fmt.Printf("For key v the value is %v \n", value)
	}
}
