package service

import (
	"context"
	"github.com/b3liv3r/users-for-gym/modules/user/models"
	"github.com/b3liv3r/users-for-gym/modules/user/repository"
	"go.uber.org/zap"
)

type UserService struct {
	repo repository.UsererRepository
	log  *zap.Logger
}

func NewUserService(repo repository.UsererRepository, log *zap.Logger) Userer {
	return &UserService{repo: repo, log: log}
}

func (s *UserService) Create(ctx context.Context, userId int, username, phone, email string) (string, error) {
	err := s.repo.Create(ctx, userId, username, phone, email)
	if err != nil {
		s.log.Error("Failed to create user", zap.Error(err))
		return "", err
	}
	return "User created successfully", nil
}

func (s *UserService) Profile(ctx context.Context, userId int) (models.User, error) {
	user, err := s.repo.Profile(ctx, userId)
	if err != nil {
		s.log.Error("Failed to get user profile", zap.Error(err))
	}
	return user, err
}

func (s *UserService) Update(ctx context.Context, userId int, username, phone, email string) (string, error) {
	err := s.repo.Update(ctx, userId, username, phone, email)
	if err != nil {
		s.log.Error("Failed to update user", zap.Error(err))
		return "", err
	}
	return "User updated successfully", nil
}

func (s *UserService) UpdateSubLvl(ctx context.Context, userId int, level int) (string, error) {
	err := s.repo.UpdateSubLvl(ctx, userId, level)
	if err != nil {
		s.log.Error("Failed to update user subscription level", zap.Error(err))
		return "", err
	}
	return "User subscription level updated successfully", nil
}

func (s *UserService) UpdateGymID(ctx context.Context, userId int, gymId int) (string, error) {
	err := s.repo.UpdateGymID(ctx, userId, gymId)
	if err != nil {
		s.log.Error("Failed to update user's current gym ID", zap.Error(err))
		return "", err
	}
	return "User current gym ID updated successfully", nil
}
