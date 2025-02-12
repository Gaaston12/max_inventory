package main

import (
	"context"
	"max_inventory/database"
	"max_inventory/settings"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Options(
			fx.Provide(
				context.Background,
				settings.New,
				database.New,
			),
			fx.Invoke(),
		))

	app.Run()
}
