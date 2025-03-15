package dto

type AddTask struct {
	Title       string `json:"title" validate:"required"`
	Description *string
	Status      *string `json:"status" validate:"oneof=new in_progress done"`
}
type UpdateTask struct {
	Title       string `json:"title" validate:"required"`
	Description *string
	Status      string `json:"status" validate:"required,oneof=new in_progress done"`
}
