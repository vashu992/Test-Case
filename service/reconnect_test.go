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

func generateSampleGetReconnect() model.GetReconnect {
	return model.GetReconnect{
		Type:        "DummyType",
		Mdn:         "1234567890",           // Dummy MDN value
		Esn:         "12345678901234567890", // Dummy ESN value
		Imei:        "123456789012345",      // Dummy IMEI value
		Plan:        "BasicPlan",            // Dummy plan value
		Zip:         "12345",                // Dummy ZIP code
		BillingCode: "BC1234",               // Dummy billing code
		StateCode:   "SC",                   // Dummy state code
		E911ADDRESS: generateMockE911ADDRESS(),
	}
}

func TestGetReconnect(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetReconnectRequest
		expectedResponse *model.GetReconnectResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetReconnectRequest{
				Session:      createMockSession(),
				StateCode:    "12312",
				GetReconnect: generateSampleGetReconnect(),
			},
			expectedResponse: &model.GetReconnectResponse{
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
					Timestamp string `xml:"timestamp,attr"`
					Status    string `xml:"status,attr"`
					Type      string `xml:"type,attr"`
					Mdn       string `xml:"mdn"`
					Esn       string `xml:"esn"`
					Warning   string `xml:"warning"`
				}{Timestamp: "20200611114406", Status: "success", Type: "WirelessHotlineResponse", Mdn: "1234567890", Esn: "54665665656", Warning: "warning"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetReconnectRequest{
				Session:      createMockSession(),
				StateCode:    "12312",
				GetReconnect: generateSampleGetReconnect(),
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
			request: &model.GetReconnectRequest{
				Session:      createMockSession(),
				StateCode:    "12312",
				GetReconnect: generateSampleGetReconnect(),
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
			request:          &model.GetReconnectRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetReconnectRequest.GetReconnect.E911ADDRESS.STREET1' Error:Field validation for 'STREET1' failed on the 'required' tag\nKey: 'GetReconnectRequest.GetReconnect.E911ADDRESS.CITY' Error:Field validation for 'CITY' failed on the 'required' tag\nKey: 'GetReconnectRequest.GetReconnect.E911ADDRESS.STATE' Error:Field validation for 'STATE' failed on the 'required' tag\nKey: 'GetReconnectRequest.GetReconnect.E911ADDRESS.ZIP' Error:Field validation for 'ZIP' failed on the 'required' tag"),
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

			response, err := client.GetReconnect(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
