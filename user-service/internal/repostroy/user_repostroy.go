package repostroy

import (
	"errors"
	"user-service/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *model.User) (*model.User, error) {
    if err := r.DB.Create(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}


func (r *UserRepository) UpdateUserFields(id uint, firstname, lastname, role string) (*model.User, error) {
    if err := r.DB.Model(&model.User{}).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "firstname": firstname,
            "lastname":  lastname,
            "role":      role,
        }).Error; err != nil {
        return nil, err
    }

    var updated model.User
    if err := r.DB.First(&updated, id).Error; err != nil {
        return nil, err
    }

    return &updated, nil
}


func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
    var user model.User

    if err := r.DB.First(&user, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // topilmadi
        }
        return nil, err
    }

    return &user, nil
}

func (r *UserRepository) ListUsers() ([]model.User, error) {
    var users []model.User

    if err := r.DB.Find(&users).Error; err != nil {
        return nil, err
    }

    return users, nil
}



// GetUserByTelegramID returns a user by telegram_id
func (r *UserRepository) GetUserByTelegramID(telegramID int64) (*model.User, error) {
    var user model.User

    if err := r.DB.Where("telegram_id = ?", telegramID).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // user topilmadi
        }
        return nil, err
    }

    return &user, nil
}

// GetTelegramIDsByUserIDs returns telegram_id list for given user IDs
func (r *UserRepository) GetTelegramIDsByUserIDs(userIDs []uint) ([]int64, error) {
    var telegramIDs []int64

    if len(userIDs) == 0 {
        return telegramIDs, nil // bo'sh slice qaytarish
    }

    // faqat telegram_id larni olish
    if err := r.DB.Model(&model.User{}).
        Where("id IN ?", userIDs).
        Pluck("telegram_id", &telegramIDs).Error; err != nil {
        return nil, err
    }

    return telegramIDs, nil
}
