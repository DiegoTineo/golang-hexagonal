package ports

import "github.com/DiegoTineo/golang-live-chat/internal/features/pets/domain"

type PetService interface {
	GetPetById(id string) (*domain.Pet, error)
	GetAllPets() ([]*domain.Pet, error)
	CreatePet(pet *domain.Pet) error
	UpdatePet(pet *domain.Pet) error
	DeletePet(id string) error
}