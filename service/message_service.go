package service

import "goSiteProject/model"

var messageStorage = make([]model.Message, 0)

func GetMesaages() []model.Message {
	return messageStorage
}

func AddMesaage(newMessage model.Message) {
	messageStorage = append(messageStorage, newMessage)
}
