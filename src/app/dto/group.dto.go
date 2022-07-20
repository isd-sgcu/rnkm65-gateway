package dto

type GroupDto struct {
	ID       string      `json:"id" validate:"uuid_optional"`
	LeaderID string      `json:"leader_id" validate:"required"`
	Token    string      `json:"token" validate:"required"`
	Members  []*UserInfo `json:"members" validate:"required"`
}

type UserInfo struct {
	ID        string `json:"id" validate:"uuid_optional"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	ImageUrl  string `json:"image_url" validate:"required"`
}

type SelectBaan struct {
	Baans []string `json:"baans" validate:"required"`
}
