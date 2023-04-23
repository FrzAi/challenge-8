package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// type Employee struct {
// 	ID       int    `json:"id" db:"id"`
// 	Fullname string `json:"full_name" db:"full_name"`
// 	Email    string `json:"email" db:"email"`
// 	Age      int    `json:"age" db:"age"`
// 	Division string `json:"division" db:"division"`
// }

type Employee struct {
	ID        int       `json:"id" gorm:"primaryKey;type:serial"`
	Fullname  string    `json:"full_name" gorm:"type:varchar(50)"`
	Email     string    `json:"email" gorm:"unique"`
	Age       int       `json:"age"`
	Division  string    `json:"division" gorm:"type:varchar(20)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (e Employee) Validation() error { // custom validation
	return validation.ValidateStruct(&e,
		validation.Field(&e.Fullname, validation.Required),
		validation.Field(&e.Email, validation.Required),
		validation.Field(&e.Age, validation.Required),
		validation.Field(&e.Division, validation.Required))
}
