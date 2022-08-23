package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// Name is the unique identifier for user in database
	Name string `json:"name,omitempty"`

	// Password is the plain text password
	Password string `json:"password,omitempty"`

	// PasswordHash is the bcrypt hashed password
	PasswordHash string `json:"-"`

	// Token is the JSON Web Token
	Token string `json:"token,omitempty"`

	// TokenId is the identifier of JSON Web Token
	TokenId string `json:"-"`

	// CreatedAt is the unix timestamp in seconds
	// indicating the time when the user was created
	CreatedAt int64 `json:"created_at,omitempty"`
}

func New(name string) (*User, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	return &User{
		Name: name,
	}, nil
}

func (u *User) PrepareNew() error {
	err := u.verifyFields()
	if err != nil {
		return err
	}
	err = u.createPasswordHash()
	if err != nil {
		return err
	}
	u.Password = ""
	u.CreatedAt = time.Now().Unix()
	return nil
}

func (u User) PrepareLogin() error {
	return u.verifyFields()
}

func (u *User) createPasswordHash() error {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(u.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *User) VerifyPassword() error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(u.PasswordHash),
		[]byte(u.Password),
	)
	if err != nil {
		return err
	}
	u.Password = ""
	return nil
}

func (u User) verifyFields() error {
	// name
	if u.Name == "" {
		return errors.New("name is empty")
	}
	if len(u.Name) > 32 {
		return errors.New("name cannot be longer than 32 characters")
	}

	// password
	if u.Password == "" {
		return errors.New("password is empty")
	}
	if len(u.Password) < 8 {
		return errors.New("password should be at least 8 characters long")
	}

	return nil
}
