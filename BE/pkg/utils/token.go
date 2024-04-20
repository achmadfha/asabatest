package utils

import (
	"BE-shop/models/dto"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var configData dto.ConfigData

func InitConfigData(data dto.ConfigData) {
	configData = data
}

func GenerateToken(id uuid.UUID, role string) (tokenString string, err error) {

	secret := configData.DbConfig.SecretToken
	expired := configData.DbConfig.TokenExpire

	fmt.Println(expired)
	now := time.Now()
	expiredTime := time.Now().Add(time.Duration(expired) * time.Hour)

	claims := jwt.MapClaims{
		"clientId": id,
		"exp":      expiredTime.Unix(),
		"role":     role,
		"iat":      now.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
