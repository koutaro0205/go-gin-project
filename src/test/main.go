package main

import (
	"fmt"
	handler "go-gin-project/test/example"
)

func main() {
	person := handler.Person{Name: "こうたろう", Age: 23}
	fmt.Println(person.Greeting("Hello!!"))
}
