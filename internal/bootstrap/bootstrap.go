package bootstrap

import (
	"context"
	"fmt"

	"github.com/jsovalles/stori_transaction_summary/internal/api"
	"github.com/jsovalles/stori_transaction_summary/internal/controller"
	"go.uber.org/fx"
)

var Module = fx.Options(
	controller.Module,
	api.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle, routes api.Api,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("--------------- Starting Story Transaction Summary Service ---------------")

			go func() {
				routes.SetupRoutes()
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("--------------- Stopping Story Transaction Summary Service ---------------")
			return nil
		},
	})
}