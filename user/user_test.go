package user_test

import (
	"testing"

	"pingmaster/user"

	"golang.org/x/crypto/bcrypt"
)

type newUser struct {
	user.User
	ExpectError bool
}

var users = []newUser{}

func init() {
	users = []newUser{
		{
			User: user.User{
				Name:     "Ramesh Deo",
				Password: "Ramesh Deo",
			},
			ExpectError: false,
		},
		{
			User: user.User{
				Name:     "Ramesh",
				Password: "Ramesh",
			},
			ExpectError: true,
		},
	}
}

func TestPrepareNew(t *testing.T) {
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
