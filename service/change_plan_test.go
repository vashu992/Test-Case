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

func generateSampleGetChangePlan() model.GetChangePlan {
	return model.GetChangePlan{
		Type:        "dummy-type",
		Mdn:         "1234567890",
		Sim:         "123123123123",
		NewplanID:   "dummyNewPlanID",
		E911ADDRESS: generateMockE911ADDRESS(),
	}
}

func TestGetChangePlan(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetChangePlanRequest
		expectedResponse *model.GetChangePlanResponse
		expectedError    error
	}{
		{
			name: "Valid request",
			request: &model.GetChangePlanRequest{
				Session:       createMockSession(),
				GetChangePlan: generateSampleGetChangePlan(),
			},
			expectedResponse: &model.GetChangePlanResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{Clec: struct {
					ID string `xml:"id"`
				}{ID: "1004"}, Timestamp: "20200603041914"},
				Response: struct {
					Timestamp string "xml:\"timestamp,attr\""
					Status    string "xml:\"status,attr\""
					Type      string "xml:\"type,attr\""
					Message   string "xml:\"message\""
					Warning   string "xml:\"warning\""
				}{
					Timestamp: "20200603041914",
					Status:    "success",
					Type:      "ChangePlanResponse",
					Message:   "SUCCESS",
					Warning:   "change plan completed successfully",
				},
			},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetChangePlanRequest{
				Session:       createMockSession(),
				GetChangePlan: generateSampleGetChangePlan(),
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
			request:          &model.GetChangePlanRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetChangePlanRequest.GetChangePlan.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetChangePlanRequest.GetChangePlan.Sim' Error:Field validation for 'Sim' failed on the 'alphanum' tag\nKey: 'GetChangePlanRequest.GetChangePlan.NewplanID' Error:Field validation for 'NewplanID' failed on the 'required' tag\nKey: 'GetChangePlanRequest.GetChangePlan.E911ADDRESS.STREET1' Error:Field validation for 'STREET1' failed on the 'required' tag\nKey: 'GetChangePlanRequest.GetChangePlan.E911ADDRESS.CITY' Error:Field validation for 'CITY' failed on the 'required' tag\nKey: 'GetChangePlanRequest.GetChangePlan.E911ADDRESS.STATE' Error:Field validation for 'STATE' failed on the 'required' tag\nKey: 'GetChangePlanRequest.GetChangePlan.E911ADDRESS.ZIP' Error:Field validation for 'ZIP' failed on the 'required' tag"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()

			client := createTestClient(mockServer.URL)

			response, err := client.GetChangePlan(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {

				assert.Equal(t, tc.expectedResponse, response)
			}
		})
	}
}
