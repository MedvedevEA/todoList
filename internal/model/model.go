package model

import "time"

type Task struct {
	Id          int
	Title       string
	Description *string
	Status      string
	Created_at  *time.Time
	Updated_at  *time.Time
}
