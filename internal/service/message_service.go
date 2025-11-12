package service

import (
	"goSiteProject/internal/model"
	"goSiteProject/internal/repository"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

var logger slog.Logger

func MustInitService(mainLogger *slog.Logger) {
	logger := *mainLogger
	logger.Info("Service was initialize")
}

func GetMesaages(topic string) ([]model.Message, error) {
	rows, err := repository.AllMessagesFromTopic(topic)

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

func AddMesaage(topic string, newMessage model.Message) (bool, error) {
	return repository.AddMessage(topic, newMessage)
}
