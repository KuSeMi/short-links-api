package auth

import (
	"demo/internal/user"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Login(email, password string) (string, error) {
	user, _ := service.UserRepository.FindByEmail(email)
	if user == nil {
		return "", errors.New(ErrInvalidCredentials)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrInvalidCredentials)
	}

	return user.Email, nil
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, err := service.UserRepository.FindByEmail(email)
	if err == nil && existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &user.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}

	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}
