package model

type PayloadNotificationRequest struct {
	Device string `json:"device" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
	Type  string `json:"type" validate:"required"`
}