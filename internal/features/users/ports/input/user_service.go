package ports

import (
	"context"

	"github.com/DiegoTineo/golang-live-chat/internal/features/users/domain"
)

// Operaciones que se pueden realizar con un usuario (casos de uso)
type UserService interface {
	RegisterUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id string) (*domain.User, error)
	UpdateUser(ctx context.Context, id string, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) error
	Login(ctx context.Context, email, password string) (*domain.User, error)
}
