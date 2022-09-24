package migration

type NotificationTable struct {
}

func (NotificationTable) TableName() string {
	return "notification"
}