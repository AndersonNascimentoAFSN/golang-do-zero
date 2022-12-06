package main

import (
	"fmt"
	"log"

	"github.com/AndersonNascimentoAFSN/golang-do-zero/crawler"
	"github.com/AndersonNascimentoAFSN/golang-do-zero/math"
	"github.com/AndersonNascimentoAFSN/golang-do-zero/ponteiro"
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

	result, err := crawler.GetUrl("http://google.com.br")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result.Status)

	sumTotal := math.SumTotal(1, 2, 3, 4)
	fmt.Println("sumTotal", sumTotal)

	//// função anônima
	// resultado := func(x ...int) func() int {
	// 	res := 0

	// 	for _, v := range x {
	// 		res += v
	// 	}

	// 	return func() int {
	// 		return res * res
	// 	}
	// }
	// fmt.Println(resultado(1, 2, 3)())
	car1 := Car{Name: "BMW", Year: 2022}
	car1.Walk()

	ponteiro.Point()
}

type Car struct {
	Name string
	Year int
}

func (c Car) Walk() {
	fmt.Println(c.Name, "Andou!")
}
