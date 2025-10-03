package user

import "gorm.io/gorm"

type IUserRepository interface {
	Save(user *User) (*User, error)
	FindByAccountId(accountId string) (*User, error)
	FindById(id uint) (*User, error)
	FindAll(limit, offset int) ([]User, error)
	DeleteById(id uint) (bool, error)
}

var _ IUserRepository = &UserRepository{}

type UserRepository struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Save(user *User) (*User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteById(id uint) (bool, error) {
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *UserRepository) FindByAccountId(accountId string) (*User, error) {
	var user User
	if err := r.db.Where("account_id=?", accountId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindById(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindAll(limit, offset int) ([]User, error) {
	var list []User
	if err := r.db.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
