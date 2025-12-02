package services

import (
	"context"

	"github.com/furqanrupom/sqlc-crud/model"
	"github.com/furqanrupom/sqlc-crud/repo"
)

type UserService interface {
	Create(ctx context.Context, user model.User) (*model.RUser, error)
	Get(ctx context.Context, id string) (*model.RUser, error)
	List(ctx context.Context) ([]*model.RUser, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, user model.RUser) (*model.RUser, error)
}

type userService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(ctx context.Context, user model.User) (*model.RUser, error) {
	return s.repo.Create(ctx,user)
}
func (s *userService) Get(ctx context.Context, id string) (*model.RUser, error)  {
	return s.repo.Get(ctx,id)
}
func (s *userService) List(ctx context.Context) ([]*model.RUser, error){
	return s.repo.List(ctx)
}
func (s *userService) Delete(ctx context.Context, id string)error{
	return s.repo.Delete(ctx,id)
}
func (s *userService) Update(ctx context.Context, user model.RUser) (*model.RUser, error) {
	return s.repo.Update(ctx,user)
}
