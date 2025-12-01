package service

import (
	"fmt"
	"time"
	"user-service/internal/model"
	"user-service/internal/repostroy"
	"user-service/internal/util/redis"
	"user-service/internal/util/validation"
)

// DTO lar (Create va Update uchun)
type CreateUserDTO struct {
	TelegramID int64
	Firstname  string
	Lastname   string
	Role       string
}

type UpdateUserDTO struct {
	ID        uint
	Firstname string
	Lastname  string
	Role      string
}

// =====================
//        SERVICE
// =====================
type UserService struct {
	repo *repostroy.UserRepository
}

// Constructor
func NewUserService(repo *repostroy.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// =====================
//       CREATE USER
// =====================
func (s *UserService) CreateUser(dto *CreateUserDTO) (*model.User, error) {
	// ✅ validation
	if err := validation.ValidateCreateUser(dto.Firstname, dto.Lastname, dto.Role, dto.TelegramID); err != nil {
		return nil, err
	}

	user := &model.User{
		TelegramID: dto.TelegramID,
		Firstname:  dto.Firstname,
		Lastname:   dto.Lastname,
		Role:       dto.Role,
	}

	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// ✅ Redis write-through
	key := fmt.Sprintf("user:%d", createdUser.ID)
	_ = redis.Set(key, createdUser, time.Minute*5)

	return createdUser, nil
}

// =====================
//       UPDATE USER
// =====================
func (s *UserService) UpdateUser(dto *UpdateUserDTO) (*model.User, error) {
	// ✅ validation
	if err := validation.ValidateUpdateUser(dto.ID, dto.Firstname, dto.Lastname, dto.Role); err != nil {
		return nil, err
	}

	updatedUser, err := s.repo.UpdateUserFields(dto.ID, dto.Firstname, dto.Lastname, dto.Role)
	if err != nil {
		return nil, err
	}

	// ✅ Redis write-through
	key := fmt.Sprintf("user:%d", updatedUser.ID)
	_ = redis.Set(key, updatedUser, time.Minute*5)

	return updatedUser, nil
}

// =====================
//      GET USER BY ID (read-through Redis)
// =====================
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	key := fmt.Sprintf("user:%d", id)

	var user model.User
	err := redis.Get(key, &user, func() (interface{}, error) {
		// fallback → DB
		return s.repo.GetUserByID(id)
	}, time.Minute*5)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// =====================
//        LIST USERS
// =====================
func (s *UserService) ListUsers() ([]model.User, error) {
	return s.repo.ListUsers()
}

// =====================
// GET USER BY TELEGRAM ID
// =====================
func (s *UserService) GetUserByTelegramID(telegramID int64) (*model.User, error) {
	if err := validation.ValidateTelegramID(telegramID); err != nil {
		return nil, err
	}
	return s.repo.GetUserByTelegramID(telegramID)
}

// =====================
//  GET TELEGRAM IDS BY USER IDS
// =====================
func (s *UserService) GetTelegramIDsByUserIDs(userIDs []uint) ([]int64, error) {
	return s.repo.GetTelegramIDsByUserIDs(userIDs)
}
