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

func generateSampleGetUpdateAddress() model.GetUpdateAddress {
	return model.GetUpdateAddress{
		Type:        "DummyType",
		Mdn:         "1234567890",
		Esn:         "ESN1234567890",
		E911Address: generateMockE911Address(),
	}
}

func TestGetUpdateAddress(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetUpdateAddressRequest
		expectedResponse *model.GetUpdateAddressResponse
		expectedError    error
	}{
		{
			name: "Valid request",
			request: &model.GetUpdateAddressRequest{
				Session:          createMockSession(),
				GetUpdateAddress: generateSampleGetUpdateAddress(),
			},
			expectedResponse: &model.GetUpdateAddressResponse{
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
					Errors    struct {
						Text  string `xml:",chardata"`
						Error struct {
							Code    string `xml:"code"`
							Message string `xml:"message"`
						}
					}
					Mdn string `xml:"mdn"`
					Sim string `xml:"sim"`
				}{
					Timestamp: "20200611114406",
					Status:    "success",
					Errors: struct {
						Text  string `xml:",chardata"`
						Error struct {
							Code    string `xml:"code"`
							Message string `xml:"message"`
						}
					}{
						Text: "Sample",
						Error: struct {
							Code    string `xml:"code"`
							Message string `xml:"message"`
						}{
							Code:    "123",
							Message: "Dummy message",
						},
					},
					Mdn: "655656236565",
					Sim: "6565653265",
				}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetUpdateAddressRequest{
				Session:          createMockSession(),
				GetUpdateAddress: generateSampleGetUpdateAddress(),
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
			request: &model.GetUpdateAddressRequest{
				Session:          createMockSession(),
				GetUpdateAddress: generateSampleGetUpdateAddress(),
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
			request:          &model.GetUpdateAddressRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetUpdateAddressRequest.GetUpdateAddress.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetUpdateAddressRequest.GetUpdateAddress.E911Address.Street1' Error:Field validation for 'Street1' failed on the 'required' tag\nKey: 'GetUpdateAddressRequest.GetUpdateAddress.E911Address.City' Error:Field validation for 'City' failed on the 'required' tag\nKey: 'GetUpdateAddressRequest.GetUpdateAddress.E911Address.State' Error:Field validation for 'State' failed on the 'required' tag\nKey: 'GetUpdateAddressRequest.GetUpdateAddress.E911Address.Zip' Error:Field validation for 'Zip' failed on the 'required' tag"),
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

			response, err := client.GetUpdateAddress(context.TODO(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
