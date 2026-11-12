package service

import (
	"fmt"
	"inventory-api/internal/dto"
	"inventory-api/internal/model"
	"inventory-api/internal/repository"
	"inventory-api/pkg/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
}


func NewAuthService(r *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: r}
}

func (s *AuthService) Register(req dto.RegisterRequest) (*model.User, error) {
	_, err := s.userRepo.GetUserByEmail(req.Email)
	if err == nil {
		return nil, fmt.Errorf("User with email %s already exists", req.Email)
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("Error hashing password: %v", err)
	}

	user := &model.User{
		Name: req.Name,
		Email: req.Email,
		Password: hashedPassword,
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("Error creating user: %v", err)
	}

	return user, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (jwt string, err error) {
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return "", fmt.Errorf("User with email %s not found", req.Email)
	}

	isMatch, err := utils.CheckPasswordHash(req.Password, user.Password)
	if err != nil {
		return "", fmt.Errorf("Error checking password: %v", err)
	}
	if !isMatch {
		return "", fmt.Errorf("Invalid password for user %s", req.Email)
	}

	return utils.GenerateJWT(user.ID, user.Role)
}