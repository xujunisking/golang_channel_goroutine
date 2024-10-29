package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {

	var personChan chan Person = make(chan Person, 10)

	for i := 1; i <= 10; i++ {
		age := rand.Intn(100)
		person := Person{"person" + strconv.Itoa(i), age, "Address" + strconv.Itoa(i)}
		personChan <- person
	}

	for i := 1; i < 10; i++ {
		person := <-personChan
		fmt.Println(person)
	}

}
