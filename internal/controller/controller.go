package controller

import (
	"database/sql"
	"errors"
	"strconv"
	ctlDto "todolist/internal/controller/dto"
	"todolist/internal/model"
	repoDto "todolist/internal/repository/dto"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	AddTask(dto *repoDto.AddTask) (*model.Task, error)
	GetTasks(dto *repoDto.GetTasks) ([]*model.Task, error)
	UpdateTask(dto *repoDto.UpdateTask) (*model.Task, error)
	RemoveTask(dto *repoDto.RemoveTask) error
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
	validate := validator.New()
	body := new(ctlDto.AddTask)
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	bodyStatus := "new"
	if body.Status == nil {
		body.Status = &bodyStatus
	}
	if err := validate.Struct(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	res, err := c.service.AddTask(&repoDto.AddTask{
		Title:       body.Title,
		Description: body.Description,
		Status:      *body.Status,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(res)

}
func (c *Controller) GetTasks(ctx *fiber.Ctx) error {
	res, err := c.service.GetTasks(&repoDto.GetTasks{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(res)
}
func (c *Controller) UpdateTask(ctx *fiber.Ctx) error {
	validate := validator.New()
	body := new(ctlDto.UpdateTask)
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := validate.Struct(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	paramId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.service.UpdateTask(&repoDto.UpdateTask{
		Id:          paramId,
		Title:       body.Title,
		Description: body.Description,
		Status:      body.Status,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(res)
}
func (c *Controller) RemoveTask(ctx *fiber.Ctx) error {
	paramId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = c.service.RemoveTask(&repoDto.RemoveTask{
		Id: paramId,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
