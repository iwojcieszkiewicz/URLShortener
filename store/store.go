package store

import (
	"errors"
	"math/rand"
	"strings"
)

type Store struct {
	urls map[string]string
}

func New() *Store {
	return &Store{
		urls: make(map[string]string),
	}
}

func (s *Store) Save(URL string) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"

	var code strings.Builder

	for range 6 {
		index := rand.Intn(len(chars))
		char := string(rune(chars[index]))

		code.WriteString(char)
	}

	s.urls[code.String()] = URL

	return code.String()
}

func (s *Store) Get(code string) (string, error) {

	url, ok := s.urls[code]

	if !ok {
		return "", errors.New("nie znaleziono")
	}

	return url, nil
}
