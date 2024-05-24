package pwg_test

import (
	"context"
	"errors"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockGetActivate() model.GetActivate {
	return model.GetActivate{
		Esn:         "89012607230000000",
		PlanId:      "dummy-plan-id",
		Language:    "dummy-language",
		Zip:         "12345", // Validating for 5 characters
		BillingCode: "billing",
		E911Address: generateMockE911Address(),
	}
}
func TestGetActivate(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetActivateRequest
		expectedResponse *model.GetActivateResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetActivateRequest{
				Session:     createMockSession(),
				GetActivate: mockGetActivate(),
			},
			expectedResponse: &model.GetActivateResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20200611114406"}, Response: struct {
					Status         string `xml:"status,attr"`
					Timestamp      string `xml:"timestamp,attr"`
					Type           string `xml:"type,attr"`
					Mdn            string `xml:"mdn"`
					Esn            string `xml:"esn"`
					CustomerID     string `xml:"CustomerID"`
					SubscriptionID string `xml:"SubscriptionID"`
					Warning        string `xml:"warning"`
				}{Status: "success", Timestamp: "20150609060644", Type: "WirelessActivateResponse", Mdn: "1234567890", Esn: "89012607230000000", CustomerID: "111111", SubscriptionID: "111111111111111", Warning: "activation completed successfully"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetActivateRequest{
				Session:     createMockSession(),
				GetActivate: mockGetActivate(),
			},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
				Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
				Err:     errors.New("EOF"),
			},
		},
		{
			name:             inValidCase,
			request:          &model.GetActivateRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetActivateRequest.GetActivate.Esn' Error:Field validation for 'Esn' failed on the 'required' tag\nKey: 'GetActivateRequest.GetActivate.PlanId' Error:Field validation for 'PlanId' failed on the 'required' tag\nKey: 'GetActivateRequest.GetActivate.Zip' Error:Field validation for 'Zip' failed on the 'required' tag\nKey: 'GetActivateRequest.GetActivate.BillingCode' Error:Field validation for 'BillingCode' failed on the 'min' tag\nKey: 'GetActivateRequest.GetActivate.E911Address.Street1' Error:Field validation for 'Street1' failed on the 'required' tag\nKey: 'GetActivateRequest.GetActivate.E911Address.City' Error:Field validation for 'City' failed on the 'required' tag\nKey: 'GetActivateRequest.GetActivate.E911Address.State' Error:Field validation for 'State' failed on the 'required' tag\nKey: 'GetActivateRequest.GetActivate.E911Address.Zip' Error:Field validation for 'Zip' failed on the 'required' tag"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()
			client := createTestClient(mockServer.URL)
			response, err := client.GetActivate(context.Background(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
