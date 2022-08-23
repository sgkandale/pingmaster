package user_test

import (
	"testing"

	"pingmaster/user"

	"golang.org/x/crypto/bcrypt"
)

type prepareNewUser struct {
	user.User
	ExpectError bool
}

func TestPrepareNew(t *testing.T) {
	users := []prepareNewUser{
		{
			User: user.User{
				Name:     "John Doe",
				Password: "John Doe",
			},
			ExpectError: false,
		},
		{
			User: user.User{
				Name:     "John",
				Password: "John",
			},
			ExpectError: true,
		},
	}

	for i, eachUser := range users {
		err := eachUser.PrepareNew()
		if err != nil && !eachUser.ExpectError {
			t.Errorf(
				"TestPrepareNew, At index %d : %s",
				i, err,
			)
		}
	}
}

func TestVerifyPassword(t *testing.T) {
	users := []user.User{
		{
			Name:     "John Doe",
			Password: "John Doe",
		},
		{
			Name:     "Jane Doe",
			Password: "Jane Doe",
		},
	}

	// generate password
	for i, eachUser := range users {
		passwordHash, err := bcrypt.GenerateFromPassword(
			[]byte(eachUser.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			t.Errorf(
				"TestVerifyPassword, At index %d : %s",
				i, err,
			)
			continue
		}
		// does not work with eachUser.PasswordHash
		users[i].PasswordHash = string(passwordHash)
	}

	// verify password
	for i, eachUser := range users {
		err := eachUser.VerifyPassword()
		if err != nil {
			t.Errorf(
				"TestVerifyPassword, At index %d : %s",
				i, err,
			)
			continue
		}
	}
}
