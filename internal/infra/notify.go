package infra

type NotificationRegister interface {
	Notify()
}

type NotificationService struct {
}
