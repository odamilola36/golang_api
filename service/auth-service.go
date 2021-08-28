package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/odamilola36/golang_api/dto"
	"github.com/odamilola36/golang_api/entity"
	"github.com/odamilola36/golang_api/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface{
	VerifyCredentials(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail (email string) bool
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRep repositories.UserRepository) AuthService {
	return &authService{userRepository: userRep}
}

func (service *authService) VerifyCredentials(email string, password string) interface{}{
	res := service.userRepository.VerifyCredentials(email, password)
	if v, ok := res.(entity.User); ok {
		comparePassword := comparePassword(v.Password, [] byte(password))
		if v.Email == email && comparePassword {
			return true
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User{
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))

	if err != nil {
		log.Println(err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) entity.User{
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool{
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPassword string, password []byte) bool {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
