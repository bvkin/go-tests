package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Stub store struct
type StubStore struct {
	response string
	cancelled bool
}
func (s *StubStore) Fetch() string {
	return s.response
}
func (s *StubStore) Cancel() {
	s.cancelled = true
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}
