package helper

import (
	"github.com/api/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SingKey = []byte("AllYourBase")

type CustomClaims struct {
	ID       uint   `json:"id" form:"id"`
	GoogleID string `json:"google_id" form:"google_id"`
	Email    string `json:"Email" form:"Email"`
	Role     string `json:"role" form:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(borrowers *models.Borrowers) (string, error) {
	claims := &CustomClaims{
		borrowers.ID,
		borrowers.Google_ID,
		borrowers.Email,
		borrowers.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Go_Api",
			Subject:   "borrowers",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SingKey)
}
