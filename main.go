package main

import (
	"max_inventory/settings"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Options(
			fx.Provide(
				settings.New,
			),
			fx.Invoke(),
		))

	app.Run()
}
