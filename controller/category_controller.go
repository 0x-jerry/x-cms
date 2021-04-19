package controller

import (
	"github.com/cwxyz007/x-cms/service"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

type CategoryController struct {
	Ctx      iris.Context
	Services *service.AllServices
	Logger   *logrus.Logger
}

func (c *CategoryController) Get() Response {
	categories, err := c.Services.Category.GetAll()

	return ResponseWithError(categories, err)
}

func (c *CategoryController) GetBy(id string) Response {
	category, err := c.Services.Category.GetBy(id)

	return ResponseWithError(category, err)
}
