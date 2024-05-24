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

func generateSampleGetCancelPurchase() model.GetCancelPurchase {
	return model.GetCancelPurchase{
		Type:       "dummy-type",
		Mdn:        "1234567890",
		PurchaseId: "dummyPurchaseId1234567890",
	}
}
func TestGetCancelPurchase(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetCancelPurchaseRequest
		expectedResponse *model.GetCancelPurchaseResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetCancelPurchaseRequest{
				Session:           createMockSession(),
				GetCancelPurchase: generateSampleGetCancelPurchase(),
			},
			expectedResponse: &model.GetCancelPurchaseResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004",
					},
					Timestamp: "20150427101032",
				},
				Response: struct {
					Timestamp  string `xml:"timestamp,attr"`
					Status     string `xml:"status,attr"`
					Type       string `xml:"type,attr"`
					PurchaseId string `xml:"purchaseId"`
				}{Timestamp: "20150427101032", Status: "success", Type: "WirelessCancelPurchaseResponse", PurchaseId: "2045"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetCancelPurchaseRequest{
				Session:           createMockSession(),
				GetCancelPurchase: generateSampleGetCancelPurchase(),
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
			request:          &model.GetCancelPurchaseRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetCancelPurchaseRequest.GetCancelPurchase.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetCancelPurchaseRequest.GetCancelPurchase.PurchaseId' Error:Field validation for 'PurchaseId' failed on the 'required' tag"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()
			client := createTestClient(mockServer.URL)
			response, err := client.GetCancelPurchase(context.Background(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				return
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
