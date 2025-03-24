package ports

import "github.com/DiegoTineo/golang-live-chat/internal/features/pets/domain"

type PetRepository interface {
	GetByID(id string) (*domain.Pet, error)
	GetAll() ([]*domain.Pet, error)
	Create(pet *domain.Pet) error
	Update(pet *domain.Pet) error
	Delete(id string) error
}