package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"required"`
	Username string	`json:"username"  validate:"required"`
	Password string	`json:"password"  validate:"required"`
	Salt     string	
	RoleName string	`json:"role_name"`
	Role     Role `gorm:"foreignKey:RoleName;references:Name"`
}

type JWTClaim struct {
	Username     string `json:"username"`
	Role         string `json:"role"`
	RandomString []byte `json:"random_string"`
	jwt.RegisteredClaims
}