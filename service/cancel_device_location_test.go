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

func generateSampleGetCancelDeviceLocation() model.GetCancelDeviceLocation {
	return model.GetCancelDeviceLocation{
		Type: "dummy-type",
		Mdn:  "1234567890",
		Esn:  "dummyESN123456789012345",
	}
}

func TestGetCancelDeviceLocation(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetCancelDeviceLocationRequest
		expectedResponse *model.GetCancelDeviceLocationResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetCancelDeviceLocationRequest{

				Session:                 createMockSession(),
				GetCancelDeviceLocation: generateSampleGetCancelDeviceLocation(),
			},
			expectedResponse: &model.GetCancelDeviceLocationResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20170627030530"},
				Response: struct {
					Timestamp string `xml:"timestamp,attr"`
					Status    string `xml:"status,attr"`
					Type      string `xml:"type,attr"`
					Message   string `xml:"message"`
					MDN       string `xml:"MDN"`
					ESN       string `xml:"ESN"`
				}{
					Timestamp: "20170627030530", Status: "success", Type: "CancelDeviceLocationResponse", Message: "SUCCESS", MDN: "4257727633", ESN: "8901260713165529250"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetCancelDeviceLocationRequest{

				Session:                 createMockSession(),
				GetCancelDeviceLocation: generateSampleGetCancelDeviceLocation(),
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
			request: &model.GetCancelDeviceLocationRequest{
				Session:                 createMockSession(),
				GetCancelDeviceLocation: generateSampleGetCancelDeviceLocation(),
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
			request:          &model.GetCancelDeviceLocationRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetCancelDeviceLocationRequest.GetCancelDeviceLocation.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag"),
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
			response, err := client.GetCancelDeviceLocation(context.Background(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
