package main

import (
	"clean-arch/domain/usecases"
	"clean-arch/entrypoints"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func ProvideRouter(app *fiber.App) fiber.Router {
	return app
}

func Run(app *fiber.App, routes *entrypoints.Routes) error {
	routes.Mount()

	return app.Listen(":3000")
}

func main() {
	app := fx.New(
		fx.Provide(
			ProvideRouter,
			entrypoints.NewRoutes,
		),
		fx.Provide(
			usecases.NewCartService,
			entrypoints.NewCartRoute,
		),
		fx.Supply(
			fiber.New(),
		),
		fx.Invoke(Run),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
