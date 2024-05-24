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

func TestGetCancelPortIn(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetCancelPortInRequest
		expectedResponse *model.GetCancelPortInResponse
		expectedError    error
	}{
		{
			name:    "Valid request",
			code:    0,
			request: &model.GetCancelPortInRequest{},
			expectedResponse: &model.GetCancelPortInResponse{ResponseSession: struct {
				Clec struct {
					ID string `xml:"id"`
				}
				Timestamp string `xml:"timestamp"`
			}{Clec: struct {
				ID string `xml:"id"`
			}{ID: "1004"}, Timestamp: "20150609123908"}, Response: struct {
				Timestamp   string "xml:\"timestamp,attr\""
				Status      string "xml:\"status,attr\""
				Type        string "xml:\"type,attr\""
				Mdn         string "xml:\"mdn\" validate:\"required,len=10,numeric\""
				Sim         string "xml:\"sim\" validate:\"required,max=25\""
				Result      string "xml:\"result\""
				Description string "xml:\"description\""
			}{
				Timestamp:   "20150609123908",
				Status:      "success",
				Type:        "CancelPortInResponse",
				Mdn:         "4052740222",
				Sim:         "8901260723154403222",
				Result:      "SUCCESS",
				Description: "PortIn request cancelled successfully",
			}},
			expectedError: nil,
		},
		{
			name:             xmlUnmarshalError,
			request:          &model.GetCancelPortInRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
				Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
				Err:     errors.New("EOF"),
			},
		},
		{
			name:             internalServerError,
			request:          &model.GetCancelPortInRequest{},
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
			response, err := client.GetCancelPortIn(context.TODO(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {

				assert.Equal(t, tc.expectedResponse, response)
			}
		})
	}
}
