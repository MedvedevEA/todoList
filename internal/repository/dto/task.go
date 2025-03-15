package dto

import "time"

type AddTask struct {
	Title       string
	Description *string
	Status      string
}
type GetTasks struct {
}
type UpdateTask struct {
	TaskId      int
	Title       string
	Description *string
	Status      string
	Created_at  *time.Time
	Updated_at  *time.Time
}
type RemoveTask struct {
	TaskId int
}
