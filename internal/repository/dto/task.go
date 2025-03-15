package dto

type AddTask struct {
	Title       string  `db:"title"`
	Description *string `db:"description"`
	Status      string  `db:"status"`
}
type GetTasks struct {
}
type UpdateTask struct {
	Id          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Status      string  `json:"status" db:"status"`
}
type RemoveTask struct {
	Id int `db:"id"`
}
