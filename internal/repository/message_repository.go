package repository

import (
	"database/sql"
	"fmt"
	"goSiteProject/internal/model"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

var logger slog.Logger

func initConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/topics.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MustCheckConnection(mainLogger *slog.Logger) {
	logger = *mainLogger
	db, err := initConnection()
	if err != nil {
		panic(err)
	}
	tableExists(db)
	db.Close()
}

func tableExists(db *sql.DB) {
	query := "SELECT name FROM sqlite_master"
	var name string
	row, _ := db.Query(query)
	for row.Next() {
		row.Scan(&name)
		logger.Info(name)
	}
}

func AllMessagesFromTopic(topic string) (*sql.Rows, error) {
	db, err := initConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM `%s`", topic))
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func AddMessage(topic string, newMessage model.Message) (bool, error) {
	db, err := initConnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name, message, date) VALUES ('%s', '%s', '%s');", topic, newMessage.Name, newMessage.Message, newMessage.Date))

	if err != nil {
		return false, err
	}

	return true, nil
}
