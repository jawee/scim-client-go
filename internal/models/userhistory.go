package models

import "time"

type UserHistory struct {
    UserName      string    `json:"userName,omitempty"`
    ErrorMessages []string  `json:"errorMessages,omitempty"`
    LastSync      time.Time `json:"lastSync,omitempty"`
}
