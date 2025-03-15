package service

import (
	"todolist/internal/model"
	"todolist/internal/repository"
	"todolist/internal/repository/dto"
)

type Service struct {
	store repository.Repository
}

func New(store repository.Repository) *Service {
	return &Service{
		store: store,
	}
}
func (s *Service) AddTask(dto *dto.AddTask) (*model.Task, error) {
	return s.store.AddTask(dto)
}
func (s *Service) GetTasks(dto *dto.GetTasks) ([]*model.Task, error) {
	return s.store.GetTasks(dto)
}
func (s *Service) UpdateTask(dto *dto.UpdateTask) (*model.Task, error) {
	return s.store.UpdateTask(dto)
}
func (s *Service) RemoveTask(dto *dto.RemoveTask) error {
	return s.store.RemoveTask(dto)
}
