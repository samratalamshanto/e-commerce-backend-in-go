package user

import "ecom-backend/internal/common"

type UserService struct {
	repo *UserRepository
}

func GetUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(dto CreateUserDTO) (*User, error) {
	//todo: can check validation
	var user *User
	user, err := common.Mapper[CreateUserDTO, User](dto)
	if err != nil {
		return nil, err
	}
	return s.repo.Save(user)
}

func (s *UserService) UpdateUser(dto UpdateUserDTO, id uint) (*User, error) {
	//todo: can check validation
	user, err := common.Mapper[UpdateUserDTO, User](dto)
	if err != nil {
		return nil, err
	}
	return s.repo.Save(user)
}

func (s *UserService) GetUserByAccountID(accountID string) (*User, error) {
	return s.repo.FindByAccountID(accountID)
}

func (s *UserService) GetUserByID(id uint) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.FindAll()
}

func (s *UserService) DeleteUserByID(id uint) (bool, error) {
	return s.repo.DeleteByID(id)
}
