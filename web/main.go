package main

import (
	"lottery-study/bootstrap"
	"lottery-study/web/middleware/identity"
	"lottery-study/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("lottery", "duxing")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.ConfigureRoutes)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8081")
	//app.Run(iris.Addr(":8080"))
}
