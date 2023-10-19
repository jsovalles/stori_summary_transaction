package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEnv(t *testing.T) {
	cfg := NewEnv()

	assert.Equal(t, "local", cfg.Environment)
	assert.Equal(t, "smtp.gmail.com", cfg.SmtpServer)
	assert.Equal(t, "587", cfg.SmtpPort)
	assert.Equal(t, "storitest0@gmail.com", cfg.SmtpUsername)
    assert.Equal(t, "fjli pxvq fdwu hojr", cfg.SmtpPassword)
    assert.Equal(t, "storitest0@gmail.com", cfg.GeneralEmail)
    assert.Equal(t, "db", cfg.DbHost)
    assert.Equal(t, "5432", cfg.DbPort)
    assert.Equal(t, "root", cfg.DbUsername)
    assert.Equal(t, "root", cfg.DbPassword)
    assert.Equal(t, "pgdb", cfg.DbSchema)
}