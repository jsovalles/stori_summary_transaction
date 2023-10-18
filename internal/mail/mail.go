package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"

	"github.com/jsovalles/stori_transaction_summary/internal/config"
	"github.com/jsovalles/stori_transaction_summary/internal/models"
	"github.com/jsovalles/stori_transaction_summary/internal/utils"
	"go.uber.org/fx"
)

const (
	EMAIL_SUBJECT = "Financial Report"
	MESSAGE_INFO  = "From: %s\nTo: %s\nSubject: %s\nMIME-Version: 1.0\nContent-Type: text/html; charset=utf-8\n\n%s"
)

type Email interface {
	SendEmail(ts models.TransactionSummary) (err error)
}

type email struct {
	config config.Config
}

func (e *email) SendEmail(ts models.TransactionSummary) (err error) {

	body, err := generateEmailBody(ts)
	if err != nil {
		return
	}

	from := mail.Address{Address: e.config.GeneralEmail}
	to := mail.Address{Address: e.config.GeneralEmail}
	msg := fmt.Sprintf(MESSAGE_INFO, from.Address, to.Address, EMAIL_SUBJECT, body)

	auth := smtp.PlainAuth("", e.config.SmtpUsername, e.config.SmtpPassword, e.config.SmtpServer)
	err = smtp.SendMail(fmt.Sprintf("%s:%s", e.config.SmtpServer, e.config.SmtpPort), auth, from.Address, []string{to.Address}, []byte(msg))
	if err != nil {
		return fmt.Errorf(utils.EmailErr, err)
	}

	return
}

func generateEmailBody(ts models.TransactionSummary) (body string, err error) {
	data := struct {
		Summary models.TransactionSummary
	}{
		Summary: ts,
	}

	tmpl, err := template.New("emailTemplate").Parse(utils.EmailTemplate)
	if err != nil {
		return "", fmt.Errorf(utils.ParsingTemplateErr, err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", fmt.Errorf(utils.ParsingTemplateErr, err)
	}

	body = buf.String()

	return
}

func NewTokenCreator(config config.Config) Email {
	return &email{config: config}
}

var Module = fx.Provide(NewTokenCreator)
