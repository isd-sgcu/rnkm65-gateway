package dto

type CheckinVerifyRequest struct {
	EventType int `json:"event_type"`
}

type CheckinConfirmRequest struct {
	Token string `json:"token"`
}
