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

func TestGetQueryUsage(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetQueryUsageRequest
		expectedResponse *model.GetQueryUsageResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetQueryUsageRequest{
				Session:   createMockSession(),
				GetMdnEsn: generateSampleGetMdnEsn(),
			},
			expectedResponse: &model.GetQueryUsageResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{ID: "1004"}, Timestamp: "20200611114406"},
				Response: struct {
					Status        string `xml:"status,attr"`
					Timestamp     string `xml:"timestamp,attr"`
					Type          string `xml:"type,attr"`
					MDN           string `xml:"MDN" validate:"required,len=10,numeric"`
					SIM           string `xml:"SIM"`
					ACCOUNTSTATUS string `xml:"ACCOUNTSTATUS"`
					SOC           string `xml:"SOC"`
					TYPE          string `xml:"TYPE"`
					LIMIT         string `xml:"LIMIT"`
					USED          string `xml:"USED"`
					USAGESTATUS   string `xml:"USAGESTATUS"`
					STATUSCODE    string `xml:"STATUSCODE"`
					DESCRIPTION   string `xml:"DESCRIPTION"`
				}{Timestamp: "20200611114406", Status: "success", Type: "QueryUsageResponse", MDN: "1234567890", SIM: "54665665656", ACCOUNTSTATUS: "Active", SOC: "SOC", LIMIT: "100", USED: "YES", USAGESTATUS: "Active", STATUSCODE: "200", DESCRIPTION: "Dummy description"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetQueryUsageRequest{
				Session:   createMockSession(),
				GetMdnEsn: generateSampleGetMdnEsn(),
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
			request:          &model.GetQueryUsageRequest{},
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

			response, err := client.GetQueryUsage(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
