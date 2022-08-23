package user

import (
	"errors"
	"fmt"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const (
	token_IssuerField     = "iss"
	token_SubjectField    = "sub"
	token_ExpiryField     = "exp"
	token_NotBeforeField  = "nbf"
	token_IssuedAtField   = "iat"
	token_IdentifierField = "jti"
	token_NameField       = "name"
)

// CreateToken creates a JSON Web Token from the user data using HMAC algorithm
// with tokenSecret as the signing key
func (u *User) CreateToken(tokenSecret []byte) error {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			token_IssuerField:     "pingmaster",
			token_SubjectField:    "access to pingmaster",
			token_ExpiryField:     time.Now().Add(time.Hour).Unix(),
			token_NotBeforeField:  time.Now().Unix(),
			token_IssuedAtField:   time.Now().Unix(),
			token_IdentifierField: uuid.New().String(),
			token_NameField:       u.Name,
		},
	)

	tokenString, err := token.SignedString(tokenSecret)
	if err != nil {
		return err
	}

	u.Token = tokenString
	return nil
}

// DecodeToken decoded the given token string into user
// with the given token secret as signing key
func DecodeToken(tokenString string, tokenSecret []byte) (*User, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return tokenSecret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		name, ok := claims[token_NameField].(string)
		if !ok {
			log.Println("[ERROR] user.DecodeToken, name field is not of string type")
			return nil, errors.New("invalid token")
		}

		return &User{
			Name: name,
		}, nil
	}

	return nil, errors.New("invalid token")
}
