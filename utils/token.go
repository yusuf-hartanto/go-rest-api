package utils

import (
	"errors"
	"rest-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("signature-key")

type JWTClaim struct {
	UserID int64 `json:"id"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (string, error) {
	expire := time.Now().Add(24 * time.Hour)

	jwtClaim := &JWTClaim{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	tokenString, err := token.SignedString(JwtKey)
	return tokenString, err
}

func ValidationToken(tokenSign string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenSign, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token expired")
	}

	return claims, err
}
