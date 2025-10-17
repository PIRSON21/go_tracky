package models

import "github.com/google/uuid"

type Board struct {
	Id         uuid.UUID `json:"id"`
	User_id    uuid.UUID `json:"user_id"`
	Name_board string    `json:"name_board"`
	Access     string    `json:"access"`
	Color      string    `json:"color"`
}
