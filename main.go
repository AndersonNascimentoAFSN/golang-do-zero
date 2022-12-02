package main

import (
	"fmt"
	"log"

	"github.com/AndersonNascimentoAFSN/golang-do-zero/crawler"
	"github.com/AndersonNascimentoAFSN/golang-do-zero/math"
)

func main() {
	//// variables
	// var name string
	// name = "Anderson Nascimento"
	// age, city := 31, "Maceió"
	// state := "Alagoas"

	// fmt.Printf("%v\n%v\n%v\n%v\n", name, age, city, state)
	// fmt.Printf("%T\n%T\n%T\n%T\n", name, age, city, state)

	total, err := math.Sum(10, 2)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%v\n", total)

	result, err := crawler.GetUrl("http://google.comr")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result.Status)
}
