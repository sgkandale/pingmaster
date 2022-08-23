package user_test

import (
	"testing"

	"pingmaster/user"
)

const (
	testTokenSecret = "test1233"
)

func TestToken(t *testing.T) {
	usr := user.User{
		Name: "John",
	}

	err := usr.CreateToken([]byte(testTokenSecret))
	if err != nil {
		t.Errorf("TestToken.CreateToken : %s", err)
		return
	}

	if usr.Token == "" {
		t.Error("TestToken : token is blank")
		return
	}

	newUsr, err := user.DecodeToken(usr.Token, []byte(testTokenSecret))
	if err != nil {
		t.Errorf("TestToken.DecodeToken : %s", err)
		return
	}
	if newUsr.Name != usr.Name {
		t.Errorf(
			"TestToken : name not matching after decoding, previous name : %s, new name : %s",
			usr.Name, newUsr.Name,
		)
	}
}
