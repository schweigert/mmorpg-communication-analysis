package forms

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/krakenlab/ternary"
)

// JwtCredentialForm can turn a token, and load data from
type JwtCredentialForm struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Token from this form
func (form *JwtCredentialForm) Token(jwtKey []byte) string {
	form.StandardClaims.ExpiresAt = time.Now().Add(5 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, form)
	tokenString, err := token.SignedString(jwtKey)
	ternary.Func(err != nil, func() { panic(err) }, func() {})()

	return tokenString
}

// Parse from jwtPayload
func (form *JwtCredentialForm) Parse(jwtKey []byte, payload string) bool {
	tkn, err := jwt.ParseWithClaims(payload, form, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return err == nil && tkn.Valid
}
