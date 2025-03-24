package main

import (
	"fmt"

	"github.com/DiegoTineo/golang-live-chat/internal/app"
	"github.com/DiegoTineo/golang-live-chat/internal/features/pets/domain"
)



func main() {

	deps := app.Bootstrap()

	err := deps.PetService.CreatePet(
		&domain.Pet{
			Id: "1",
			Name: "Firulais",
			PetType: "Dog",
			PetSound: "Guau",
		},
	)

	if err != nil {
		fmt.Println("Error creating pet", err)
	} else {
		fmt.Println("Pet created")
	}

	pet, err := deps.PetService.GetPetById("1")
	
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Print(pet)
	}


}