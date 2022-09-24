package model

type PayloadNotificationRequest struct {
	Device string `json:"device"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}

type PayloadNotificationArtikel struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Type  string `json:"type"`
}

type PayloadNotificationKabarDonasi struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	Type     string `json:"type"`
	DonasiId string `json:"donasi_id"`
}

type ResponseNotificationBlast struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ResponseNotification struct {
	Device string `json:"device"`
	Title  string `json:"title"`
	Type   string `json:"type"`
}
