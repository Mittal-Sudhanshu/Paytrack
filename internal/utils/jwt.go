package utils

import (
	"time"

	"github.com/example/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key") // Replace with os.Getenv("JWT_SECRET") in production

func GenerateJWT(user entity.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role_id": user.RoleId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
