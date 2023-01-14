package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := `I was all alone with the love of my life
	She's got glitter for skin
	My radiant beam in the night
	I don't need no light to see you`

	file, err := os.Create("./mylyrics.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkErrNil(err)
	length, err := io.WriteString(file, content)
	checkErrNil(err)

	fmt.Println("The length is", length)

	defer file.Close()

	readFile("./mylyrics.txt")
}

func readFile(fileName string) {
	datatype, err := ioutil.ReadFile(fileName)
	checkErrNil(err)
	fmt.Println("The content is", string(datatype))
}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}