package auth

import (
	"os"
	"time"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/golang-jwt/jwt/v5"
)

const (
	sourceJWT = "Generate JWT f(n) under auth pkg"
)

func (a *Auth) generateJWT(claim data.ActiveUser) ([]byte, error) {
	// Step one: fill in the struct values of the claim.RegisteredClaims type declared in the ActiveUser struct
	claim.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "http://localhost:8080",
		Subject:   "Access token",
		Audience:  jwt.ClaimStrings{claim.Username},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1 * time.Hour))),
	}

	// Step two: Generate the token using the new with claims function
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtKey := os.Getenv("JWTKEY")

	// Step three: sign the token using the jwt key stored in the .env file and the environments tag of docker compose file
	signedToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		e := helper.WrapError("Token signing error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceJWT,
		})
		return nil, e
	}

	// Step four: return the signed token
	return []byte(signedToken), nil
}
