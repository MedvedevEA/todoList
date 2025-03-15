package dto

type AddTask struct {
	Title       string `json:"title" validate:"required,min=0"`
	Description *string
	Status      string
}
type UpdateTask struct {
	Title       string
	Description *string
	Status      string
}
