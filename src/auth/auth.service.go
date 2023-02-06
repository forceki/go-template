package auth

import (
	"log"
	"time"

	"github.com/forceki/invent-be/config"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UsersRes struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthService interface {
	Create(Data UsersRes) error
	Login(Username string, Password string) (interface{}, error)
}

type authService struct {
	authrepository AuthRepository
}

func NewAuthService(authrepository AuthRepository) *authService {
	return &authService{authrepository}
}

func (s *authService) Create(Data UsersRes) error {

	var user Users

	hash, err := bcrypt.GenerateFromPassword([]byte(Data.Password), 10)
	if err != nil {
		return err
	}

	log.Println(Data)

	user.Email = Data.Email
	user.Username = Data.Username
	user.Phone = Data.Phone
	user.Password = string(hash)

	err = s.authrepository.Create(user)

	return err
}

func (s *authService) Login(Username string, Password string) (interface{}, error) {
	var user Users

	user, err := s.authrepository.FindUsername(Username)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))

	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"user_id":  user.Id,
		"username": user.Username,
		"email":    user.Email,
		"phone":    user.Phone,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	secret := config.Config("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return nil, err
	}

	type response struct {
		Users UsersRes `json:"user"`
		Token string   `json:"token"`
	}

	res := response{}

	res.Users.Username = user.Username
	res.Users.Email = user.Email
	res.Users.Phone = user.Phone
	res.Token = t

	return res, err
}
