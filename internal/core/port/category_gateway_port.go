package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type CategoryGateway interface {
	FindByID(ctx context.Context, id uint64) (*entity.Category, error)
	FindAll(ctx context.Context, name string, page, limit int) ([]*entity.Category, int64, error)
	Create(ctx context.Context, category *entity.Category) error
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id uint64) error
}
