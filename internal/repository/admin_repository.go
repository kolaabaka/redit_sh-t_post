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
	db.Close()
	return true, nil
}

func CheckUserCreds(login string, password string) int {
	db, err := initConnection()
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256([]byte(password))
	row, err := db.Query("SELECT id FROM \"users\" WHERE login = ? AND pas_hash = ?;", login, fmt.Sprintf("%x", hash))
	if err != nil || !row.Next() {
		row.Close()
		return -1
	}
	var id int
	row.Scan(&id)
	row.Close()
	db.Close()
	return id
}
