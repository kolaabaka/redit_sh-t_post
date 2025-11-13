package repository

import (
	"crypto/sha256"
	"fmt"
)

func AddUser(login string, password string) (bool, error) {
	db, err := initConnection()
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256([]byte(password))

	_, err = db.Exec("INSERT INTO \"users\"(login, pas_hash) VALUES(?, ?);", login, fmt.Sprintf("%x", hash))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
