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

func generateSampleGetAddWFC() model.GetAddWFC {
	return model.GetAddWFC{
		Type:        "dummy-type",
		Mdn:         "1234567890",
		Esn:         "dummy-esn",
		E911ADDRESS: generateMockE911ADDRESS(),
	}
}

func TestGetAddWFC(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetAddWFCRequest
		expectedResponse *model.GetAddWFCResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetAddWFCRequest{

				Session:   createMockSession(),
				GetAddWFC: generateSampleGetAddWFC(),
			},
			expectedResponse: &model.GetAddWFCResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20200603093905"}},
						
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetAddWFCRequest{

				Session:   createMockSession(),
				GetAddWFC: generateSampleGetAddWFC(),
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
			request: &model.GetAddWFCRequest{

				Session:   createMockSession(),
				GetAddWFC: generateSampleGetAddWFC(),
			},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_FAILED_CODE,
				Message: constants.REQUEST_FAILED_INFO,
				Err:     errors.New("Post \"\": unsupported protocol scheme \"\""),
			},
		},
		{
			name:             inValidCase,
			request:          &model.GetAddWFCRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetAddWFCRequest.GetAddWFC.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetAddWFCRequest.GetAddWFC.E911ADDRESS.STREET1' Error:Field validation for 'STREET1' failed on the 'required' tag\nKey: 'GetAddWFCRequest.GetAddWFC.E911ADDRESS.CITY' Error:Field validation for 'CITY' failed on the 'required' tag\nKey: 'GetAddWFCRequest.GetAddWFC.E911ADDRESS.STATE' Error:Field validation for 'STATE' failed on the 'required' tag\nKey: 'GetAddWFCRequest.GetAddWFC.E911ADDRESS.ZIP' Error:Field validation for 'ZIP' failed on the 'required' tag"),
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

			response, err := client.GetAddWFC(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
