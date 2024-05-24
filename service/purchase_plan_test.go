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

func generateSample() model.GetPurchasePlan {
	return model.GetPurchasePlan{
		Type:        "dummyType",
		Mdn:         "1234567890",
		PlanId:      "dummyPlanId",
		E911ADDRESS: generateMockE911ADDRESS(),
	}
}

func TestGetPurchasePlan(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetPurchasePlanRequest
		expectedResponse *model.GetPurchasePlanResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetPurchasePlanRequest{
				Session:         createMockSession(),
				GetPurchasePlan: generateSample(),
			},
			expectedResponse: &model.GetPurchasePlanResponse{
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
					Timestamp  string `xml:"timestamp,attr"`
					Status     string `xml:"status,attr"`
					Type       string `xml:"type,attr"`
					PurchaseId string `xml:"purchaseId"`
					Warning    string `xml:"warning"`
				}{
					Timestamp:  "20200611114406",
					Status:     "success",
					Type:       "QueryPortStatusResponse",
					PurchaseId: "545454564564",
					Warning:    "dfdfdf",
				}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetPurchasePlanRequest{
				Session:         createMockSession(),
				GetPurchasePlan: generateSample(),
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
			request: &model.GetPurchasePlanRequest{
				Session:         createMockSession(),
				GetPurchasePlan: generateSample(),
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
			request:          &model.GetPurchasePlanRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetPurchasePlanRequest.GetPurchasePlan.E911ADDRESS.STREET1' Error:Field validation for 'STREET1' failed on the 'required' tag\nKey: 'GetPurchasePlanRequest.GetPurchasePlan.E911ADDRESS.CITY' Error:Field validation for 'CITY' failed on the 'required' tag\nKey: 'GetPurchasePlanRequest.GetPurchasePlan.E911ADDRESS.STATE' Error:Field validation for 'STATE' failed on the 'required' tag\nKey: 'GetPurchasePlanRequest.GetPurchasePlan.E911ADDRESS.ZIP' Error:Field validation for 'ZIP' failed on the 'required' tag"),
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

			response, err := client.GetPurchasePlan(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
