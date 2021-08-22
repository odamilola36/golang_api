package dto

type UserUpdateDTO struct {
	ID 			uint64 `json:"id" form:"id" binding:"required"`
	Name 		string `json:"name" form:"name" binding:"required"`
	Email 		string `json:"email" form:"email" binding:"required" validate:"email"`
	Password 	string `json:"password" form:"password,omitempty" validate:"min:8"`
}

type CreateUserDTO struct {
	Name 		string `json:"name" form:"name" binding:"required"`
	Email 		string `json:"email" form:"email" binding:"required" validate:"email"`
	Password 	string `json:"password" form:"password,omitempty" validate:"min:8" binding:"required"`
}
