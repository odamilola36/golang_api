package repositories

import (
	"log"

	"github.com/odamilola36/golang_api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//userepository is a contract that defines actions that can be takenn on the db
type UserRepository interface {
	InsertUser(user entity.User)
	UpdateUser(user entity.User)
	VerifyCredentials(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userId string) entity.User
}

//gives access to the db
type userConnection struct {
	connection *gorm.DB
}

//creates a new instance of UserRepository
//throwing this error because I've not used this NewUserRepository method anywhere
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{connection: db}
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("failed to hash password")
	}
	return string(hash)
}

func(db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func(db *userConnection) UpdateUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return entity.User{}
}

func(db *userConnection) VerifyCredentials(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return nil
	}
	return user;
}

func(db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func(db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func(db *userConnection) ProfileUser(userId string) entity.User {
	var user entity.User
	db.connection.Find(&user, userId)
	return user 
}