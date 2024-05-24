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

func generateSampleGetChangePWGCostPlan() model.GetChangePWGCostPlan {
	return model.GetChangePWGCostPlan{
		Type:        "dummy-type",
		Mdn:         "1234567890",
		Billingcode: "dummyBillingCode",
	}
}

func TestGetChangePWGCostPlan(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetChangePWGCostPlanRequest
		expectedResponse *model.GetChangePWGCostPlanResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetChangePWGCostPlanRequest{

				Session:              createMockSession(),
				GetChangePWGCostPlan: generateSampleGetChangePWGCostPlan(),
			},
			expectedResponse: &model.GetChangePWGCostPlanResponse{
				Credentials: struct {
					ReferenceNumber string `xml:"referenceNumber"`
					ReturnURL       string `xml:"returnURL"`
				}{ReferenceNumber: "", ReturnURL: ""}, WholeSaleOrderResponse: struct {
					StatusCode  string `xml:"statusCode"`
					Description string `xml:"description"`
				}{StatusCode: "0", Description: "INTER"}}, expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetChangePWGCostPlanRequest{

				Session:              createMockSession(),
				GetChangePWGCostPlan: generateSampleGetChangePWGCostPlan(),
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
			request: &model.GetChangePWGCostPlanRequest{

				Session:              createMockSession(),
				GetChangePWGCostPlan: generateSampleGetChangePWGCostPlan(),
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
			request:          &model.GetChangePWGCostPlanRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetChangePWGCostPlanRequest.GetChangePWGCostPlan.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetChangePWGCostPlanRequest.GetChangePWGCostPlan.Billingcode' Error:Field validation for 'Billingcode' failed on the 'required' tag"),
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

			response, err := client.GetChangePWGCostPlan(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {

				assert.Equal(t, tc.expectedResponse, response)
			}
		})
	}
}
