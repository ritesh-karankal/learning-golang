package main

import "fmt"

func main() {
	ritesh := User{"Ritesh", "ritesh.dev", true, 20}
	fmt.Println(ritesh)
	fmt.Printf("The details are %+v \n", ritesh)
	fmt.Printf("The name is %v and email is %v \n", ritesh.Name, ritesh.Email)
	ritesh.GetStatus()
	ritesh.ChangeEmail()

	fmt.Printf("email is %v \n", ritesh.Email)
	
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("User is active:", u.Status)
}

func (u User) ChangeEmail() {
	u.Email = "test.dev"
	fmt.Println("The email is", u.Email)
}