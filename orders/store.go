package main

import (
	"context"
)

type Store struct {
	//db *sql.DB
}

func NewStore() *Store {
	return &Store{}
}

/*func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}*/

func (s *Store) Create(ctx context.Context) error {
	return nil
}
