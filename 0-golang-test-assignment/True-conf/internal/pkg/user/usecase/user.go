package usecase

import (
	"userService/internal/pkg/models"
	"userService/internal/pkg/user/repository"
)

type UserUsecaseInterface interface {
	CreateUser(user models.User) (models.User, error)
	GetUserById(id int) (models.User, error)
	DeleteUser(id int) error
	ChangeUser(user models.User, id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

type UserUsecase struct {
	Repository repository.UserRepositoryInterface
}

func (u UserUsecase) GetAllUsers() ([]models.User, error) {
	return u.Repository.GetAllUsers()
}

func (u UserUsecase) ChangeUser(user models.User, id int) (models.User, error) {
	user.Id = id
	return user, u.Repository.ChangeUser(user)
}

func (u UserUsecase) DeleteUser(id int) error {
	return u.Repository.DeleteUser(id)
}

func (u UserUsecase) GetUserById(id int) (models.User, error) {
	return u.Repository.GetUserById(id)
}

func (u UserUsecase) CreateUser(user models.User) (models.User, error) {
	id, err := u.Repository.CreateUser(user)

	if err != nil {
		return models.User{}, err
	}

	user.Id = id

	return user, nil
}
