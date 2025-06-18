package service

import (
	"context"
	"errors"
	"time"

	"github.com/example/internal/entity"
	"github.com/example/internal/model"
	"github.com/example/internal/repository"
	"github.com/example/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	ListUsers(ctx context.Context) (any, error)
	SignupUser(ctx context.Context, req model.SignupRequest) (any, error)
	SignupUserReturnUser(ctx context.Context, req model.SignupRequest) (any, error)

	Login(ctx context.Context, req model.LoginRequest) (any, error)
}

type userService struct {
	userRepo repository.Repository[entity.User]
}

func NewUserService(userRepo repository.Repository[entity.User]) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) SignupUser(ctx context.Context, req model.SignupRequest) (any, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := entity.User{
		Email:         req.Email,
		Password_hash: string(hashedPassword),
		RoleId:        req.RoleId,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		PhoneNumnber:  req.PhoneNumber,
		LastLogin:     time.Now(),
		// LastLogin can be set to zero value or time.Now() if needed
	}

	createdUser, err := s.userRepo.Create(ctx, &user)
	if err != nil {
		return nil, err
	}
	jwt, err := utils.GenerateJWT(*createdUser)
	if err != nil {
		return nil, errors.New("error generating jwt")
	}
	return jwt, nil
}
func (s *userService) SignupUserReturnUser(ctx context.Context, req model.SignupRequest) (any, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := entity.User{
		Email:         req.Email,
		Password_hash: string(hashedPassword),
		RoleId:        req.RoleId,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		PhoneNumnber:  req.PhoneNumber,
		LastLogin:     time.Now(),
		// LastLogin can be set to zero value or time.Now() if needed
	}

	createdUser, err := s.userRepo.Create(ctx, &user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
func (s *userService) Login(ctx context.Context, req model.LoginRequest) (any, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email And Password Are required")
	}

	// Replace the filter below with the appropriate filter type for your repository, e.g., map or struct
	user, err := s.userRepo.Query(ctx, map[string]interface{}{"email": req.Email})
	if err != nil {
		return nil, err
	}
	decryptedPassword := bcrypt.CompareHashAndPassword([]byte(user[len(user)-1].Password_hash), []byte(req.Password))
	if decryptedPassword != nil {
		return nil, errors.New("passwords do not match")
	}
	//update last login
	// After successful password check
	user[len(user)-1].LastLogin = time.Now()
	updateErr := s.userRepo.DB.WithContext(ctx).Model(&user).Update("last_login", user[len(user)-1].LastLogin).Error
	if updateErr != nil {
		return nil, updateErr
	}
	jwt, err := utils.GenerateJWT(user[len(user)-1])

	if err != nil {
		return nil, errors.New("error generating jwt")
	}
	return jwt, nil

}
func (s *userService) ListUsers(ctx context.Context) (any, error) {
	users, err := s.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
