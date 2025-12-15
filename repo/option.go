package repo

import (
	"context"

	"github.com/furqanrupom/sqlc-crud/db"
	"github.com/furqanrupom/sqlc-crud/model"
)

type OptionRepo interface {
	Create(ctx context.Context, payload model.Option) (*model.Option, error)
	Update(ctx context.Context, payload model.Option) (*model.Option, error)
	Delete(ctx context.Context, ID int) error
	Get(ctx context.Context) ([]*model.Option, error)
	Find(ctx context.Context, ID int) (*model.Option, error)
}

type optionRepo struct {
	db *db.Queries
}

func NewOptionRepo(db *db.Queries) OptionRepo {
	return &optionRepo{
		db: db,
	}
}

func (r *optionRepo) Create(ctx context.Context, payload model.Option) (*model.Option, error) {
	return nil, nil
}

func (r *optionRepo) Update(ctx context.Context, payload model.Option) (*model.Option, error) {
	return nil, nil
}
func (r *optionRepo) Delete(ctx context.Context, ID int) error {
	return nil
}

func (r *optionRepo) Get(ctx context.Context) ([]*model.Option, error) {
	return []*model.Option{}, nil
}

func (r *optionRepo) Find(ctx context.Context, ID int) (*model.Option, error) {
	return nil, nil
}
