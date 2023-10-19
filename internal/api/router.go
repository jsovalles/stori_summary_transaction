package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jsovalles/stori_transaction_summary/internal/controller"
	"go.uber.org/fx"
)

type Api struct {
	SummaryTransactionController controller.SummaryTransactionController
}

func (api *Api) SetupRoutes() {
	r := gin.Default()

	apiRoute := r.Group("/api")
	apiRoute.POST("/sign-up", api.SummaryTransactionController.SignUpAccount)
	apiRoute.POST("/account/:id/upload-transactions", api.SummaryTransactionController.UploadTransactions)
	apiRoute.GET("/account/:id/transactions", api.SummaryTransactionController.ListAccountTransactionsByAccountId)

	if err := r.Run(":8080"); err != nil {
		log.Fatalln("Failed to listen and serve on port 8080: " + err.Error())
		panic(err)
	}

}

func NewRoutes(summaryTransactionController controller.SummaryTransactionController) Api {
	return Api{SummaryTransactionController: summaryTransactionController}
}

var Module = fx.Provide(NewRoutes)
