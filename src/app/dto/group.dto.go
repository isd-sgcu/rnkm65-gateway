package dto

type GroupDto struct {
	ID       string `json:"id" validate:"uuid_optional"`
	LeaderID string `json:"leader_id" validate:"required"`
	Token    string `json:"token" validate:"required"`
}
