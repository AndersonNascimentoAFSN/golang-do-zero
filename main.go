package main

import (
	"fmt"

	"github.com/AndersonNascimentoAFSN/golang-do-zero/crawler"
)

func main() {
	//// variables
	// var name string
	// name = "Anderson Nascimento"
	// age, city := 31, "Maceió"
	// state := "Alagoas"

	// fmt.Printf("%v\n%v\n%v\n%v\n", name, age, city, state)
	// fmt.Printf("%T\n%T\n%T\n%T\n", name, age, city, state)

	// total := math.Sum(1, 2)
	// fmt.Printf("%T\n", total)

	result := crawler.GetUrl("http://google.com.br")
	fmt.Println(result.Status)
}
