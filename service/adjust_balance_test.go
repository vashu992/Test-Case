package pwg_test

import (
	"context"
	"errors"
	"net/http"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateSampleGetAdjustBalanceSingle() model.GetAdjustBalanceSingle {
	return model.GetAdjustBalanceSingle{
		Mdn:            "1234567890",
		SubscriptionId: "123123",
		Uom:            "dummy-uom",
		Amount:         "100",
		ExpiryDate:     "2024-12-31",
	}
}

func TestGetAdjustBalanceSingle(t *testing.T) {
	testCases := []struct {
		name             string
		request          *model.GetAdjustBalanceSingleRequest
		expectedResponse *model.GetAdjustBalanceSingleResponse
		expectedError    error
		code             int
	}{
		{
			name: validCase,
			request: &model.GetAdjustBalanceSingleRequest{
				Session:                createMockSession(),
				GetAdjustBalanceSingle: generateSampleGetAdjustBalanceSingle(),
			},
			expectedResponse: &model.GetAdjustBalanceSingleResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{Clec: struct {
					ID string `xml:"id"`
				}{ID: "1004"}, Timestamp: "20200611052857"},
				Response: struct {
					Timestamp string `xml:"timestamp,attr"`
					Status    string `xml:"status,attr"`
					Type      string `xml:"type,attr"`
					MDN       string `xml:"MDN"`
				}{Timestamp: "20200611052857", Status: "success", Type: "WirelessAdjustBalanceResponse", MDN: "1234567890"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetAdjustBalanceSingleRequest{
				Session:                createMockSession(),
				GetAdjustBalanceSingle: generateSampleGetAdjustBalanceSingle(),
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
			request:          &model.GetAdjustBalanceSingleRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetAdjustBalanceSingleRequest.GetAdjustBalanceSingle.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetAdjustBalanceSingleRequest.GetAdjustBalanceSingle.SubscriptionId' Error:Field validation for 'SubscriptionId' failed on the 'required' tag"),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()
			client := createTestClient(mockServer.URL)
			response, err := client.GetAdjustBalanceSingle(context.TODO(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}

func generateMockGetAdjustMultipleBalancePayload() model.GetAdjustMultipleBalancePayload {
	subscriptions := model.Subscriptions{
		SubscriptionId: "1234567890",
		ExpiryDate:     "2024-12-31",
		Subscription: []model.Subscription{
			{
				Text:   "Sample Text 1",
				Uom:    "unit",
				Amount: "100.00",
			},
			{
				Text:   "Sample Text 2",
				Uom:    "unit",
				Amount: "200.00",
			},
		},
	}

	payload := model.GetAdjustMultipleBalancePayload{
		Mdn:           "0123456789",
		Subscriptions: subscriptions,
	}

	return payload
}

func TestGetAdjustMultipleBalance(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetAdjustMultipleBalanceRequest
		expectedResponse *model.GetAdjustMultipleBalanceResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetAdjustMultipleBalanceRequest{

				Session: createMockSession(),
				Request: generateMockGetAdjustMultipleBalancePayload(),
			},
			expectedResponse: &model.GetAdjustMultipleBalanceResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{Clec: struct {
					ID string `xml:"id"`
				}{ID: ""}, Timestamp: ""},
				Response: struct {
					Timestamp string `xml:"timestamp,attr"`
					Status    string `xml:"status,attr"`
					Type      string `xml:"type,attr"`
					MDN       string `xml:"MDN"`
				}{Timestamp: "", Status: "", Type: "", MDN: ""}},
			expectedError: nil,
			code:          http.StatusOK,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetAdjustMultipleBalanceRequest{

				Session: createMockSession(),
				Request: generateMockGetAdjustMultipleBalancePayload(),
			},
			code:             http.StatusOK,
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
				Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
				Err:     errors.New("EOF"),
			},
		},
		{
			name:             inValidCase,
			request:          &model.GetAdjustMultipleBalanceRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetAdjustMultipleBalanceRequest.Request.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetAdjustMultipleBalanceRequest.Request.Subscriptions.SubscriptionId' Error:Field validation for 'SubscriptionId' failed on the 'required' tag"),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()
			client := createTestClient(mockServer.URL)
			response, err := client.GetAdjustMultipleBalance(context.TODO(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
