package model

type PayloadNotificationRequest struct {
	Device string `json:"device"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string `json:"user_id"`
	Type  string `json:"type"`
}

type ResponseNotification struct {
	Message string `json:"message"`
}