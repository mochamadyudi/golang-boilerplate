package utils

import (
	"errors"
	"log"
	"strconv"
	"time"

	"core.yuyuid.id/pkg/config"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userID uint) (string, error) {
	cnf := config.Get()
	var jwtSecret = []byte(cnf.Jwt.Secret)
	expiresHour, err := strconv.Atoi(cnf.Jwt.ExpiresHour)
	if err != nil {
		log.Fatal(err)
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * time.Duration(expiresHour)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ParseToken(_token string) (jwt.MapClaims, error) {
	cnf := config.Get()
	var jwtSecret = []byte(cnf.Jwt.Secret)

	token, err := jwt.Parse(_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("Invalid Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid claims")
	}
	return claims, nil
}
