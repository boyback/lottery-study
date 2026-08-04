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
func (c *IndexController) GetGifts() map[string]interface{} {
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = "success"
	result["list"] = c.ServiceGift.GetList()
	return result
}
func (c *IndexController) GetNewPrize() map[string]interface{} {
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = "success"
	return result
}
