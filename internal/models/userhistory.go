package models

import "time"

type UserHistory struct {
    UserName      string    `json:"userName,omitempty"`
    ErrorMessage  string    `json:"errorMessages,omitempty"`
    LastSync      time.Time `json:"lastSync,omitempty"`
}
