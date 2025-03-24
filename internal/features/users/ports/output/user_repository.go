package ports

import (
	"context"

	"github.com/DiegoTineo/golang-live-chat/internal/features/users/domain"
)

// Operaciones de acceso a datos de usuarios (puertos de salida, ejemplo: base de datos, cache, archivos, api, etc)
type UserRepository interface {
	
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	
}