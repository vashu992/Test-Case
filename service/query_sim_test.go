package pwg_test

import (
	"context"
	"errors"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
	pwg "propwg/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateSampleGetQuerySIM() model.GetQuerySIM {
	return model.GetQuerySIM{
		Esn: "12345678901234567890", // Dummy ESN value
	}
}

func TestGetQuerySIM(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetQuerySIMRequest
		expectedResponse *model.GetQuerySIMResponse
		expectedError    error
	}{
		{
			name: "Valid request",
			request: &model.GetQuerySIMRequest{
				Session:     createMockSession(),
				GetQuerySIM: generateSampleGetQuerySIM(),
			},
			expectedResponse: &model.GetQuerySIMResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20210105080808"},
				Response: struct {
					Status             string `xml:"status,attr"`
					Timestamp          string `xml:"timestamp,attr"`
					Type               string `xml:"type,attr"`
					SIM                string `xml:"SIM"`
					ACTIVATIONELIGIBLE string `xml:"ACTIVATIONELIGIBLE"`
					WFCEligible        string `xml:"WFCEligible"`
					PUK1               string `xml:"PUK1"`
					PUK2               string `xml:"PUK2"`
					CREATEDATE         string `xml:"CREATEDATE"`
					EXPIRATIONDATE     string `xml:"EXPIRATIONDATE"`
					ICCIDSTATUS        string `xml:"ICCIDSTATUS"`
					StatusCode         string `xml:"statusCode"`
					Description        string `xml:"description"`
				}{Timestamp: "20200611114406", Status: "success", Type: "QuerySimResponse", SIM: "1234567890", ACTIVATIONELIGIBLE: "True", WFCEligible: "True", PUK1: "puk1", PUK2: "puk2", CREATEDATE: "05/23/2024", EXPIRATIONDATE: "05/20/2025", StatusCode: "200", ICCIDSTATUS: "status", Description: "Dummy Description"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetQuerySIMRequest{
				Session:     createMockSession(),
				GetQuerySIM: generateSampleGetQuerySIM(),
			},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
				Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
				Err:     errors.New("EOF"),
			},
		},
		{
			name:             internalServerError,
			request:          &model.GetQuerySIMRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_FAILED_CODE,
				Message: constants.REQUEST_FAILED_INFO,
				Err:     errors.New("Post \"\": unsupported protocol scheme \"\""),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()

			var client *pwg.Client
			if tc.name == internalServerError {
				client = createTestClient("")
			} else {
				client = createTestClient(mockServer.URL)
			}

			response, err := client.GetQuerySIM(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
