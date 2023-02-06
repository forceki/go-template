package auth

import "gorm.io/gorm"

type Users struct {
	Id       int    `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password,omitempty"`
}

type AuthRepository interface {
	Create(Data Users) error
	FindUsername(Username string) (Users, error)
}

type authrepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authrepository {
	return &authrepository{db: db}
}

func (r *authrepository) Create(Data Users) error {
	data := Data

	err := r.db.Table("users").Create(&data).Error

	return err
}

func (r *authrepository) FindUsername(Username string) (Users, error) {
	var data Users
	err := r.db.Table("users").Select("username, email, phone, password").Find(&data).Error

	return data, err
}
