package repo

import (
	"context"

	"github.com/furqanrupom/sqlc-crud/db"
	"github.com/furqanrupom/sqlc-crud/model"
	"github.com/furqanrupom/sqlc-crud/util"
)

type UserRepo interface {
	Create(ctx context.Context, user model.User) (*model.RUser, error)
	Get(ctx context.Context, id string) (*model.RUser, error)
	List(ctx context.Context) ([]*model.RUser, error)
	Delete(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, user model.RUser) (*model.User, error)
}

type userRepo struct {
	db *db.Queries
}

func NewUserRepo(db *db.Queries) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, user model.User) (*model.RUser, error) {
	hashPass,err := util.HashPassword(user.Password,util.Params)
	if err != nil {
		return nil,err
	}
	user.Password = hashPass
	arg := db.CreateUserParams{
		Name:user.Name,
		Email: user.Email,
		Password: user.Password,
	}
	ID,err := r.db.CreateUser(ctx,arg)
	if err != nil {
		return nil, err
	}
	return &model.RUser{
		ID:int(ID),
		Name: user.Name,
		Email: user.Email,
	},nil
	
}
func (r *userRepo) List(ctx context.Context) ([]*model.RUser, error) {
    rows, err := r.db.GetUsers(ctx)
    if err != nil {
        return nil, err
    }

    users := make([]*model.RUser, len(rows))
    for i, row := range rows {
        users[i] = &model.RUser{
            ID:        int(row.ID),
            Name:      row.Name,
            Email:     row.Email,
        }
    }

    return users, nil
}

func (r *userRepo) Get(ctx context.Context, id string) (*model.RUser, error) {
	return &model.RUser{}, nil
}
func (r *userRepo) Update(ctx context.Context, user model.RUser) (*model.User, error) {
	return &model.User{}, nil
}
func (r *userRepo) Delete(ctx context.Context, id string) (*model.User, error) {
	return &model.User{}, nil
}
