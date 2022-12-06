package ponteiro

import "fmt"

func Point() {
	a := 10
	fmt.Println("ponteiro a", a, &a) // valor e endereço de memória

	var ponteiro *int = &a
	*ponteiro = 50
	fmt.Println("ponteiro", *ponteiro) // retorna o valor que esta armazenado no endereço da memória

	b := 10
	abc(&b)
	println("b", b)

	car1 := Car{
		Name: "Ka",
	}

	car1.run()
}

func abc(a *int) {
	*a = 200
}

type Car struct {
	Name string
}

func (c *Car) run() {
	c.Name = "BMW"
	fmt.Println(c.Name, "run")
}
