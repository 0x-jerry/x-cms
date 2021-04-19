package controller

import (
	"github.com/cwxyz007/x-cms/service"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

type TagController struct {
	Ctx      iris.Context
	Services *service.AllServices
	Logger   *logrus.Logger
}

func (c *TagController) Get() Response {
	tags, err := c.Services.Tag.GetAll()

	return ResponseWithError(tags, err)
}

func (c *TagController) GetBy(id string) Response {
	tag, err := c.Services.Tag.GetBy(id)

	return ResponseWithError(tag, err)
}
