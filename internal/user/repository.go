package user

import (
	"ecom-backend/internal/common"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindByAccountID(accountID string) (*User, error)
}

var _ UserRepositoryInterface = &UserRepository{}

type UserRepository struct {
	*common.BaseRepository[User]
}

func GetUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{common.GetBaseRepository[User](db)}
}

func (r *UserRepository) FindByAccountID(accountID string) (*User, error) {
	var user User
	err := r.BaseRepository.DB.Where("account_id=?", accountID).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
