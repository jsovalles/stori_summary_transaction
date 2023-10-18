package bootstrap

import (
	"context"
	"fmt"

	"github.com/jsovalles/stori_transaction_summary/internal/api"
	"github.com/jsovalles/stori_transaction_summary/internal/config"
	"github.com/jsovalles/stori_transaction_summary/internal/controller"
	"github.com/jsovalles/stori_transaction_summary/internal/mail"
	"github.com/jsovalles/stori_transaction_summary/internal/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.ConfigModule,
	mail.Module,
	service.Module,
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