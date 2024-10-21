package services

import (
	"errors"
	"fmt"
	"project1/app/helpers"
	"project1/app/models"
	"project1/app/repositories"
	"project1/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Claims struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
	jwt.RegisteredClaims
}

type AuthService interface {
	Login(email string, password string) (*models.User, error)
	CreateToken(user *models.User) (string, error)
	RefreshToken(refreshToken string) (string, error)
	ValidateToken(token string) (bool, error)
}

type AuthServiceImpl struct {
	DB      *gorm.DB
	userSvc repositories.UserRepository
}

func (c *AuthServiceImpl) Login(email string, password string) (*models.User, error) {
	user, err := c.userSvc.SearchByEmail(email)
	fmt.Println(err)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !helpers.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("password doesn't match")
	}

	return user, nil
}

func (c *AuthServiceImpl) CreateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (c *AuthServiceImpl) ValidateToken(token string) (bool, error) {
	// parts := strings.SplitN(token, " ", 2)
	// if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
	// 	return false, errors.New("Invalid Authorization Header")
	// }

	// tokenStr := parts[1]
	// // claims := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{})
	// // claims := &Claims{}

	// // error with this, already exactly same with docs
	// token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
	// 	// if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
	// 	// 	return nil, errors.New("asas")
	// 	// }
	// 	return []byte(config.Secret), nil
	// })
	// if err != nil {

	// }

	// if _, ok := token.Claims.(*Claims); ok && token.Valid {
	// 	return true, nil
	// }

	return false, errors.New("something wrong")
}

func (c *AuthServiceImpl) RefreshToken(refreshToken string) (string, error) {
	return "", errors.New("not implemented")
}

func NewAuthService(db *gorm.DB, userSvc repositories.UserRepository) AuthService {
	return &AuthServiceImpl{
		DB:      db,
		userSvc: userSvc,
	}
}
