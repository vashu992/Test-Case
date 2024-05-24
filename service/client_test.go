package pwg_test

import (
	"net/http"
	pwg "propwg/service"
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/memory"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	authConfig := pwg.AuthConfig{
		Url:      "http://example.com",
		ClecId:   "123",
		UserName: "user",
		Token:    "token",
		Pin:      "pin",
	}

	logger := &log.Logger{
		Handler: memory.New(),
	}
	httpClient := &http.Client{}

	client := pwg.NewClient(authConfig, logger, httpClient)

	assert.Equal(t, authConfig, client.AuthConfig)
	assert.Equal(t, logger, client.Logger)
	assert.Equal(t, httpClient, client.HTTPClient)
}
