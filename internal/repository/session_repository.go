package repository

import "fmt"

func AddSession(session string, id int) bool {
	db, err := initConnection()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO sessions (session_id, user_id) VALUES(?, ?);", session, id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	db.Close()
	return true
}

func RemoveSession(session string) bool {
	db, err := initConnection()
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("DELETE FROM sessions WHERE session_id = ?;", session)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
