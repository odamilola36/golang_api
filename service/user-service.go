package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/odamilola36/golang_api/dto"
	"github.com/odamilola36/golang_api/entity"
	"github.com/odamilola36/golang_api/repositories"
)


type UserService interface {
	UpdateUser(user dto.UserUpdateDTO) entity.User
	Profile(userId string) entity.User
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) UpdateUser(user dto.UserUpdateDTO) entity.User{
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil{
		log.Fatalf("Failed to map %v", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}


func (service *userService) Profile(userId string) entity.User{
	return service.userRepository.ProfileUser(userId)
}
