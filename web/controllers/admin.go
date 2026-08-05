package controllers

import (
	"lottery-study/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.GiftService
}

func (c *AdminController) Get() mvc.Result {
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "",
		},
		Layout: "admin/layout.html",
	}
}
func (c *AdminController) GetGift() mvc.Result {
	return mvc.View{
		Name: "admin/gift.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "gift",
		},
		Layout: "admin/layout.html",
	}
}
