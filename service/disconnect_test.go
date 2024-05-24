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

func generateSampleGetMdnEsn() model.GetMdnEsn {
	return model.GetMdnEsn{
		Type: "dummy-type",
		Esn:  "dummy-esn",
		Mdn:  "dummy-mdn",
	}
}

func TestGetDisconnect(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetDisconnectRequest
		expectedResponse *model.GetDisconnectResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetDisconnectRequest{

				Session:   createMockSession(),
				GetMdnEsn: generateSampleGetMdnEsn(),
			},
			expectedResponse: &model.GetDisconnectResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20150422071834"},
				Response: struct {
					Timestamp string `xml:"timestamp,attr"`
					Status    string `xml:"status,attr"`
					Type      string `xml:"type,attr"`
					Mdn       string `xml:"mdn"`
					Esn       string `xml:"esn"`
				}{Timestamp: "20150422071834", Status: "success", Type: "WirelessDisconnectResponse", Mdn: "", Esn: "89012607222222"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetDisconnectRequest{

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
			name: internalServerError,
			request: &model.GetDisconnectRequest{

				Session:   createMockSession(),
				GetMdnEsn: generateSampleGetMdnEsn(),
			},
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

			response, err := client.GetDisconnect(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {

				assert.Equal(t, tc.expectedResponse, response)
			}
		})
	}
}
