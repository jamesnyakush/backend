package main

import "github.com/nyumbapoa/backend/app/helpers"

type info struct {
	Name string
}

func main() {

	//helpers.CreateUserPesaSwap("Trial","greater","trial@gmail.com","254746445198")

	helpers.Info{
		Name: "James Nyakundi",
	}.SendEMail()
}
