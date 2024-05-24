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

func generateSampleGetDeviceSwapNotification() model.GetDeviceSwapNotification {
	return model.GetDeviceSwapNotification{
		Mdn:     "1234567890",
		Newimei: "dummy-imei",
	}
}

func TestGetDeviceSwapNotification(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetDeviceSwapNotificationRequest
		expectedResponse *model.GetDeviceSwapNotificationResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetDeviceSwapNotificationRequest{

				Session:                   createMockSession(),
				GetDeviceSwapNotification: generateSampleGetDeviceSwapNotification(),
			},
			expectedResponse: &model.GetDeviceSwapNotificationResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20220525030437"},
				Response: struct {
					Timestamp string `xml:"timestamp,attr"`
					Status    string `xml:"status,attr"`
					Type      string `xml:"type,attr"`
					Mdn       string `xml:"mdn"`
					Imei      string `xml:"imei"`
				}{Timestamp: "20220525030437", Status: "success", Type: "DeviceSwapNotificationResponse", Mdn: "5034844522", Imei: "983207423156290"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetDeviceSwapNotificationRequest{

				Session:                   createMockSession(),
				GetDeviceSwapNotification: generateSampleGetDeviceSwapNotification(),
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
			request: &model.GetDeviceSwapNotificationRequest{
				Session:                   createMockSession(),
				GetDeviceSwapNotification: generateSampleGetDeviceSwapNotification(),
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
			request:          &model.GetDeviceSwapNotificationRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetDeviceSwapNotificationRequest.GetDeviceSwapNotification.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetDeviceSwapNotificationRequest.GetDeviceSwapNotification.Newimei' Error:Field validation for 'Newimei' failed on the 'required' tag"),
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
			response, err := client.GetDeviceSwapNotification(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {

				assert.Equal(t, tc.expectedResponse, response)
			}
		})
	}
}
