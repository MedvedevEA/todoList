package repository

import (
	"todolist/internal/model"
	"todolist/internal/repository/dto"
)

type Repository interface {
	AddTask(dto *dto.AddTask) (*model.Task, error)
	GetTasks(dto *dto.GetTasks) ([]*model.Task, error)
	UpdateTask(dto *dto.UpdateTask) (*model.Task, error)
	RemoveTask(dto *dto.RemoveTask) error
}
