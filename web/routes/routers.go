package routes

import (
	"lottery-study/bootstrap"
	"lottery-study/services"
	"lottery-study/web/controllers"
	"lottery-study/web/middleware"

	"github.com/kataras/iris/v12/mvc"
)

func ConfigureRoutes(b *bootstrap.Bootstrapper) {
	giftService := services.NewGiftService()
	index := mvc.New(b.Party("/"))
	index.Register(giftService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(giftService)
	admin.Handle(new(controllers.AdminController))

	adminGift := admin.Party("/gift")
	adminGift.Router.Use(middleware.BasicAuth)
	adminGift.Register(giftService)
	adminGift.Handle(new(controllers.AdminGiftController))
}
