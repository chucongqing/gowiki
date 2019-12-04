package server

import(
	"database/sql"
)

type Store interface {
	CreateUser(user *User) error
	GetUser(name string) (*User, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateUser(user *User) error{
	
}