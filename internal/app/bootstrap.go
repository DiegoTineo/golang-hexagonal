package app

import (
	petInPorts "github.com/DiegoTineo/golang-live-chat/internal/features/pets/ports/input"
	petOutPorts "github.com/DiegoTineo/golang-live-chat/internal/features/pets/ports/output"
	petsRepository "github.com/DiegoTineo/golang-live-chat/internal/features/pets/repository"
	PetsService "github.com/DiegoTineo/golang-live-chat/internal/features/pets/service"
)

type Dependencies struct {
	// config *config.Config
	PetService petInPorts.PetService
}

func Bootstrap() *Dependencies {

	// 1. Cargar configuración
	// cfg := config.Load()

	// 2. Inicializar repositorios (Output Ports)
	var userRepo petOutPorts.PetRepository

	// if cfg.Env == "test" {
	if true {
		userRepo = petsRepository.NewInMemoryPetsRepository()
	} else {
		// otro tipo de repositorio
	}

	// 3. Inicializar servicios (Input Ports)
	appPetService := PetsService.NewPetService(userRepo)

	// 4. Retornar la capa de presentación o adaptadores
	return &Dependencies{
		PetService: appPetService,
	}
}
