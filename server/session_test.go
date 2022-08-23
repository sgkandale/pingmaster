package server_test

import (
	"testing"

	"pingmaster/server"
)

var testSessions *server.Sessions

func TestNewSessions(t *testing.T) {
	testSessions = server.NewSessions()

	if testSessions == nil {
		t.Errorf("TestNewSessions : testSessions should not be nil")
	}
}

func TestAddToken(t *testing.T) {
	testSessions.AddToken("abcd")

	tokenValid, exist := testSessions.Tokens["abcd"]
	if !exist {
		t.Error("TestAddToken : token 'abcd' should be present in testSessions")
	}
	if !tokenValid {
		t.Error("TestAddToken : token 'abcd' should be valid in testSessions")
	}
}

func TestTokenExists(t *testing.T) {
	exist := testSessions.TokenExists("abcd")
	if !exist {
		t.Error("TestAddToken : token 'abcd' should be present in testSessions")
	}
}

func TestDeleteToken(t *testing.T) {
	deleted := testSessions.DeleteToken("abcd")
	if !deleted {
		t.Error("TestAddToken : token 'abcd' should have been deleted from testSessions")
	}

	_, exist := testSessions.Tokens["abcd"]
	if exist {
		t.Error("TestAddToken : token 'abcd' should not be present in testSessions")
	}
}
