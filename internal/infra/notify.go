package infra

import "log"

type NotificationRegister interface {
	Notify(string, string,string)
}

type NotificationService struct {
	ExampleOfKafka interface{}
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (e *NotificationService) Notify(oldIP, newIP,userID string) {
	log.Printf(userID,"changed ip by",oldIP,"on",newIP)
}
