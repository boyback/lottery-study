package controllers

import (
	"lottery-study/services"

	"github.com/kataras/iris/v12"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceGift services.GiftService
}

func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "hello"
}
