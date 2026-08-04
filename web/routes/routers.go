package routes

import (
	"lottery-study/bootstrap"
	"lottery-study/services"
	"lottery-study/web/controllers"

	"github.com/kataras/iris/v12/mvc"
)

func ConfigureRoutes(b *bootstrap.Bootstrapper) {
	giftService := services.NewGiftService()
	index := mvc.New(b.Party("/"))
	index.Register(giftService)
	index.Handle(new(controllers.IndexController))
}
