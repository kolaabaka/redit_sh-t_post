package service

import (
	"database/sql"
	"fmt"
	"goSiteProject/model"

	_ "github.com/mattn/go-sqlite3"
)

func initConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/topics.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MustCheckConnection() {
	db, err := initConnection()
	if err != nil {
		panic(err)
	}
	fmt.Println("DataBase connection successfully")
	tableExists(db)
	db.Close()
}

func tableExists(db *sql.DB) {
	//todo: MAKE REPOSITORY LAYER
	query := "SELECT name FROM sqlite_master"
	var name string
	row, _ := db.Query(query)
	for row.Next() {
		row.Scan(&name)
		fmt.Println(name)
	}
}

func GetMesaages(topic string) ([]model.Message, error) {
	db, err := initConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//todo: MAKE REPOSITORY LAYER
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM `%s`", topic))

	if err != nil {
		return nil, err
	}

	var messageStorage = make([]model.Message, 0)
	for rows.Next() {
		var mesBuf model.Message
		rows.Scan(&mesBuf.Name, &mesBuf.Message, &mesBuf.Date)
		messageStorage = append(messageStorage, mesBuf)
	}

	return messageStorage, nil
}

func AddMesaage(topic string, newMessage model.Message) error {
	db, err := initConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	//todo: MAKE REPOSITORY LAYER
	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, message, date) VALUES ('%s', '%s', '%s');", topic, newMessage.Name, newMessage.Message, newMessage.Date))

	if err != nil {
		return err
	}

	return nil

}
