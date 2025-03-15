package controller

import (
	"strconv"
	"todolist/internal/model"
	"todolist/internal/repository/dto"
	"todolist/internal/servererrors"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	AddTask(dto *dto.AddTask) (*model.Task, error)
	GetTasks(dto *dto.GetTasks) ([]*model.Task, error)
	UpdateTask(dto *dto.UpdateTask) (*model.Task, error)
	RemoveTask(dto *dto.RemoveTask) error
}
type Controller struct {
	service Service
}

func Init(app *fiber.App, service Service) {
	controller := &Controller{
		service: service,
	}
	app.Post("/tasks", controller.AddTask)
	app.Get("/tasks", controller.GetTasks)
	app.Put("/tasks/:id", controller.UpdateTask)
	app.Delete("/tasks/:id", controller.RemoveTask)
}

func (c *Controller) AddTask(ctx *fiber.Ctx) error {
	req := new(dto.AddTask)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	res, err := c.service.AddTask(req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(res)
}
func (c *Controller) GetTasks(ctx *fiber.Ctx) error {
	res, err := c.service.GetTasks(&dto.GetTasks{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(res)
}
func (c *Controller) UpdateTask(ctx *fiber.Ctx) error {
	req := new(dto.UpdateTask)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	taskId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req.Id = taskId

	res, err := c.service.UpdateTask(req)
	if err == servererrors.ErrorRecordNotFound {
		return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(res)
}
func (c *Controller) RemoveTask(ctx *fiber.Ctx) error {
	taskId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = c.service.RemoveTask(&dto.RemoveTask{
		Id: taskId,
	})
	if err == servererrors.ErrorRecordNotFound {
		return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
