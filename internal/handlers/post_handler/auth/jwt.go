package auth

import (
	"time"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/golang-jwt/jwt/v5"
)

func (a *Auth) GenerateJWT(claim data.ActiveUser) {
	claim.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "http://localhost:8080",
		Subject:   "Access token",
		Audience:  jwt.ClaimStrings{claim.Username},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1 * time.Hour))),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)

	token, err := jwt.ParseWithClaims("",claim,func(t *jwt.Token) (any, error) {
		
	})
}
