package controller

import (
	"todolist/internal/model"
	repoDto "todolist/internal/repository/dto"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	AddTask(dto *repoDto.AddTask) (*model.Task, error)
	GetTasks(dto *repoDto.GetTasks) ([]*model.Task, error)
	UpdateTask(dto *repoDto.UpdateTask) (*model.Task, error)
	RemoveTask(dto *repoDto.RemoveTask) error
}

func Init(app *fiber.App, service Service) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}
