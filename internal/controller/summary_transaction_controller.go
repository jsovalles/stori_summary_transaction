package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type SummaryTransactionController interface {
	HelloWorld(ctx *gin.Context)
}

type summaryTransactionController struct {
}

func NewSummaryTransactionController() SummaryTransactionController {
	return &summaryTransactionController{
	}
}

func (c *summaryTransactionController) HelloWorld(ctx *gin.Context) {

	log.Info("hello world")

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Hello world",
	})
}



var Module = fx.Provide(NewSummaryTransactionController)