package server

import "sync"

type Sessions struct {
	// Mutex to lock the variable for concurrent access
	Mutex sync.RWMutex

	// Tokens holds the token identifiers which are currently active
	Tokens map[string]bool
}

func NewSessions() *Sessions {
	return &Sessions{
		Tokens: make(map[string]bool),
	}
}

// tokenExists checks for the existance of the token
// in sessions and returns true if it exists
func (s *Sessions) TokenExists(tokenId string) bool {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	tokenIdExist, exists := s.Tokens[tokenId]
	if !exists {
		return false
	}
	return tokenIdExist
}

// addToken adds the token to the sessions
func (s *Sessions) AddToken(tokenId string) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	s.Tokens[tokenId] = true
}

// deleteToken deletes the token from the sessions
// and returns true on deletion
func (s *Sessions) DeleteToken(tokenId string) bool {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	_, exists := s.Tokens[tokenId]
	if !exists {
		return false
	}

	delete(s.Tokens, tokenId)
	return true
}
