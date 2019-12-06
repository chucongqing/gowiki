package server

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Store interface {
	CreateUser(user *User) error
	GetUser(name string) (*User, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateUser(user *User) error {

	db, err := sql.Open("mysql", "vbi:123456@tcp(127.0.0.1:3306)/ccq")

	if err != nil {
		log.Fatal(err)
	}

	store.db = db

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
