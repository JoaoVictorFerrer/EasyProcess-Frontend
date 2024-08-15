package utils

import (
	"EasyProcess/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWTToken(userRole string, userEmail string) (string, error) {
	expirationTime := time.Now().Add(360 * time.Hour) // 15 days
	claims := &models.Claims{
		Role: userRole,
		StandardClaims: jwt.StandardClaims{
			Subject:   userEmail,
			ExpiresAt: expirationTime.Unix(),
		},
	}
	jwtKeyStringValue := os.Getenv("JWT_KEY")
	var jwtKey = []byte(jwtKeyStringValue)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWTToken(tokenString string) (claims *models.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		jwtKeyStringValue := os.Getenv("JWT_KEY")
		return []byte(jwtKeyStringValue), nil
	})
	claims, ok := token.Claims.(*models.Claims)
	if err != nil {
		return claims, err
	}
	if !ok {
		return claims, err
	}
	return claims, nil
}

func ExtractEmailFromToken(jwtToken string) (string, error) {
	claims, err := ValidateJWTToken(jwtToken)
	if err != nil {
		return "", err
	}
	return claims.Subject, nil
}
