package repository

import (
	"errors"
	"sync"

	"github.com/DiegoTineo/golang-live-chat/internal/features/pets/domain"
	ports "github.com/DiegoTineo/golang-live-chat/internal/features/pets/ports/output"
)

// mu: para garantizar que no haya problemas de concurrencia en la lectura y escritura de la lista de mascotas
// es decir, por ejemplo, que no se pueda leer la lista de mascotas mientras se está modificando

type InMemoryPetsRepository struct {
	mu   sync.RWMutex
	pets []*domain.Pet
}

func NewInMemoryPetsRepository() ports.PetRepository {
	return &InMemoryPetsRepository{
		pets: []*domain.Pet{}, // Inicializar el slice de pets (vacío)
	}
}

func (r *InMemoryPetsRepository) GetByID (id string) (*domain.Pet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, pet := range r.pets {
		if pet.Id == id {
			return pet, nil
		}
	}
	return nil, errors.New("pet not found")
}

func (r *InMemoryPetsRepository) GetAll() ([]*domain.Pet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.pets, nil
}

func (r *InMemoryPetsRepository) Create(pet *domain.Pet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.pets = append(r.pets, pet)
	return nil
}

func (r *InMemoryPetsRepository) Update(pet *domain.Pet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, p := range r.pets {
		if p.Id == pet.Id {
			r.pets[i] = pet
			return nil
		}
	}
	return errors.New("pet not found")
}

func (r *InMemoryPetsRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, pet := range r.pets {
		if pet.Id == id {
			r.pets = append(r.pets[:i], r.pets[i+1:]...)
			return nil
		}
	}
	return errors.New("pet not found")
}

