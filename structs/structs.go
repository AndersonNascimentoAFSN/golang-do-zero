package structs

import (
	"encoding/json"
	"fmt"
	"log"
)

// type ClientName string
// type ClientEmail string
// type ClientCPF int

type Client struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   int    `json:"cpf"`
}

func (c Client) Show() {
	fmt.Println("Exibindo o cliente pelo método: ", c.Name)
}

type ClientInternacional struct {
	Client
	Country string `json:"country"`
}

func HandleStruct() {
	// var ClientName ClientName
	// fmt.Println(ClientName)

	client1 := Client{
		Name:  "Anderson",
		Email: "anderson.nascimentoafsn@gmail.com",
		CPF:   12345678912,
	}
	fmt.Println(client1)

	client2 := Client{"Yanni", "yanni@gmail.com", 390810231723}
	fmt.Println(client2)

	client3 := ClientInternacional{
		Client: Client{
			Name:  "John Doe",
			Email: "john.doe@gmail.com",
			CPF:   12345678912,
		},
		Country: "USA",
	}
	fmt.Printf("Nome: %s. Email: %s. Country: %s. CPF: %d\n", client3.Name, client3.Email, client3.Country, client3.CPF)

	client2.Show()
	client3.Show()

	clientJson, err := json.Marshal(client3)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(clientJson))
	fmt.Println(client3)

	jsonClient4 := `{"name":"John Doe","email":"john.doe@gmail.com","cpf":12345678912,"country":"USA"}`
	client4 := ClientInternacional{}

	json.Unmarshal([]byte(jsonClient4), &client4)

	fmt.Println(client4)
}
