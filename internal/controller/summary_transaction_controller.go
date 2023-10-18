package controller

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jsovalles/stori_transaction_summary/internal/models"
	"github.com/jsovalles/stori_transaction_summary/internal/service"
	"github.com/jsovalles/stori_transaction_summary/internal/utils"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

const (
	FILE_EXTENSION = ".csv"
	MB             = 10 << 20
)

type SummaryTransactionController interface {
	UploadTransactions(ctx *gin.Context)
	SignUpAccount(ctx *gin.Context)
}

type summaryTransactionController struct {
	service service.SummaryTransactionService
}

func (c *summaryTransactionController) SignUpAccount(ctx *gin.Context) {
	var newAccount models.Account

	if err := ctx.BindJSON(&newAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := c.service.SignUpAccount(newAccount)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": account.Id,
	})
}

func (c *summaryTransactionController) UploadTransactions(ctx *gin.Context) {

	accountID := ctx.Request.Header.Get("account-id")
	if accountID == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.AccountIdErr})
		return
	}

	err := ctx.Request.ParseMultipartForm(MB)
	if err != nil {
		ctx.String(http.StatusBadRequest, utils.ParseFormErr)
		return
	}

	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, utils.UploadedFileErr)
		return
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	if ext != FILE_EXTENSION {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidFileErr})
		return
	}

	records, err := c.service.SummaryTransaction(file, accountID)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": records,
	})
}

func NewSummaryTransactionController(summaryTransactionService service.SummaryTransactionService) SummaryTransactionController {
	return &summaryTransactionController{
		service: summaryTransactionService,
	}
}

var Module = fx.Provide(NewSummaryTransactionController)
