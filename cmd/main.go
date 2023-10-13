package main

import (
	"github.com/jsovalles/stori_transaction_summary/internal/bootstrap"
	"go.uber.org/fx"
)

func main() {
    fx.New(bootstrap.Module).Run()
}