package entity

type Book struct {
	ID 			int64  `gorm:"primary_key:auto_increment" json:"id"`
	Title 		string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	UserId 		uint64 `gorm:"not null" json:"-"`
	User 		User   `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}