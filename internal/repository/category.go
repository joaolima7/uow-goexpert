package repository

import (
	"context"
	"database/sql"

	"github.com/joaolima7/uow-goexpert/internal/db"
	"github.com/joaolima7/uow-goexpert/internal/entity"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type CategoryRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(dtb *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (cr *CategoryRepository) Insert(ctx context.Context, category entity.Category) error {
	return cr.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: category.Name,
	})
}
