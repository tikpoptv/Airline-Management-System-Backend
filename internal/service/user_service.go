package service

import (
	"airline-management-system/internal/models"
	"airline-management-system/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	return s.repo.CreateUser(user)
}
