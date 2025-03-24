package service

import (
	"github.com/DiegoTineo/golang-live-chat/internal/features/pets/domain"
	inputPorts "github.com/DiegoTineo/golang-live-chat/internal/features/pets/ports/input"
	outputPorts "github.com/DiegoTineo/golang-live-chat/internal/features/pets/ports/output"
)

type PetService struct {
	repo outputPorts.PetRepository
}

func NewPetService(repo outputPorts.PetRepository) inputPorts.PetService {
	return &PetService{repo: repo}
}

// 	GetPetById(id string) (*domain.Pet, error)
// 	GetAllPets() ([]*domain.Pet, error)
// 	CreatePet(pet *domain.Pet) error
// 	UpdatePet(pet *domain.Pet) error
// 	DeletePet(id string) error
// 	PetSound(id string) (string, error)

func (s *PetService) GetPetById(id string) (*domain.Pet, error) {
	return s.repo.GetByID(id)
}

func (s *PetService) GetAllPets() ([]*domain.Pet, error) {
	return s.repo.GetAll()
}

func (s *PetService) CreatePet(pet *domain.Pet) error {
	return s.repo.Create(pet)
}

func (s *PetService) UpdatePet(pet *domain.Pet) error {
	return s.repo.Update(pet)
}

func (s *PetService) DeletePet(id string) error {
	return s.repo.Delete(id)
}