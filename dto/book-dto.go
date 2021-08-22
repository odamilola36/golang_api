package dto

type BookUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description"`
	UserID uint64 `json:"user_id" form:"user_id,omitempty"`
}

type BookCreateDTO struct {
	Title string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description"`
	UserID uint64 `json:"user_id" form:"user_id,omitempty"`
}