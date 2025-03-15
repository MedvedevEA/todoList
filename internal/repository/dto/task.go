package dto

type AddTask struct {
	Title       string
	Description *string
	Status      string
}
type GetTasks struct {
}
type UpdateTask struct {
	Id          int
	Title       string
	Description *string
	Status      string
}
type RemoveTask struct {
	Id int
}
