package utils

import (
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
