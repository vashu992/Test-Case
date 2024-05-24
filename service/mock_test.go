package pwg_test

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"propwg/model"
	pwg "propwg/service"

	"github.com/apex/log"
	"github.com/apex/log/handlers/memory"
)

const (
	validCase           = "Valid request"
	internalServerError = "Internal server error"
	xmlUnmarshalError   = "XML unmarshal error"
	inValidCase         = "InValid request"
)

func createMockServer(expectedResponse interface{}, statuscode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseXML, err := xml.Marshal(expectedResponse)
		if err != nil {
			fmt.Println("error while marshal response , err :", err)
		}
		if statuscode == 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(statuscode)
		}
		_, err = w.Write(responseXML)
		if err != nil {
			fmt.Println("error while writing response , err :", err)
		}

	})
	return httptest.NewServer(handler)
}

func createTestClient(serverURL string) *pwg.Client {
	authConfig := pwg.AuthConfig{
		Url:      serverURL,
		ClecId:   "123",
		UserName: "user",
		Token:    "token",
		Pin:      "pin",
	}
	logger := &log.Logger{
		Handler: memory.New(),
	}
	return &pwg.Client{
		AuthConfig: authConfig,
		Logger:     logger,
		HTTPClient: http.DefaultClient,
	}
}

func createMockSession() model.AuthSession {
	return model.AuthSession{
		AuthClec: model.AuthClec{
			ID: "dummy-id",
			AuthAgentUser: model.AuthAgentUser{
				UserName: "dummy-username",
				Token:    "dummy-token",
				Pin:      "dummy-pin",
			},
		},
	}
}
