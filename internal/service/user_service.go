package service

import (
	"errors"

	"airline-management-system/internal/models"
	"airline-management-system/internal/repository"
	"airline-management-system/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) RegisterPassenger(req *models.RegisterRequest) (*models.User, error) {
	if req.Role != "passenger" {
		return nil, errors.New("only 'passenger' role is allowed for self-registration")
	}
	return s.registerUser(req.Username, req.Email, req.Password, req.Role)
}

func (s *UserService) AdminCreateUser(req *models.AdminCreateUserRequest) (*models.User, error) {
	allowed := map[string]bool{"admin": true, "crew": true, "maintenance": true}
	if !allowed[req.Role] {
		return nil, errors.New("admin can only create admin, crew, or maintenance users")
	}
	return s.registerUser(req.Username, req.Email, req.Password, req.Role)
}

func (s *UserService) registerUser(username, email, password, role string) (*models.User, error) {
	if u, _ := s.repo.GetUserByUsername(username); u != nil && u.ID != 0 {
		return nil, errors.New("username already exists")
	}
	if u, _ := s.repo.GetUserByEmail(email); u != nil && u.ID != 0 {
		return nil, errors.New("email already exists")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Username:       username,
		Email:          email,
		HashedPassword: hashedPassword,
		Role:           role,
		IsActive:       true,
	}
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(req *models.LoginRequest) (string, *models.User, error) {
	user, err := s.repo.GetUserByUsername(req.Username)
	if err != nil || user.ID == 0 {
		return "", nil, errors.New("invalid username or password")
	}

	if err := utils.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return "", nil, errors.New("invalid username or password")
	}
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", nil, err
	}

	return signedToken, user, nil
}

func (s *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil || user.ID == 0 {
		return nil, errors.New("invalid username or password")
	}
	if err := utils.CheckPasswordHash(password, user.HashedPassword); err != nil {
		return nil, errors.New("invalid username or password")
	}
	return user, nil
}

func (s *UserService) GetUserProfile(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
