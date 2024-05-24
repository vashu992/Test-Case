package main

import (
	"accolite-gosdk/pwg"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	constants "propwg/common"
	util "propwg/common"

	"github.com/vashu992/Test-Case/model"
)

func main() {

	authConfig := pwg.AuthConfig{
		Url:      "https://oss.vcarecorporation.com:22712/api/",
		ClecId:   "12345",
		UserName: "GOSDK",
		Token:    "2312",
		Pin:      "1234",
	}

	logger := &log.Logger{}
	httpClient := &http.Client{}
	client := pwg.NewClient(authConfig, logger, httpClient)

	ctx := context.TODO()
	clientResp, err := client.GetCoverage(ctx, &model.GetCoverageRequest{CoverageRequest: model.CoverageRequest{
		Request: model.Coverage{
			Carrier: "TMB",
			Zip:     "53201",
		},
	}})
	if err != nil {
		logger.Error(err.Error())
		var wrappedError *util.WrappedError
		if ok := errors.As(err, &wrappedError); ok {
			// Handle accordingly
			switch wrappedError.Code {
			case constants.REQUEST_VALIDATION_EXCPETION_CODE:
				fmt.Print(wrappedError.Error())
			case constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE:
				fmt.Print(wrappedError.Error())
			default:
				fmt.Print(wrappedError.Error())
			}
		}
	}
	logger.WithField("res", clientResp).Info("Get Coverage response.")
}
