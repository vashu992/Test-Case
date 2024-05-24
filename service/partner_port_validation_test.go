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

func generateSampleGetPartnerPortOutValidation() model.GetPartnerPortOutValidation {
	return model.GetPartnerPortOutValidation{
		Type:        "dummyType",
		Mdn:         "1234567890",
		Sim:         "dummySim",
		MessageCode: "1234567990",
		Description: "1234567890",
	}
}

func TestGetPartnerPortOutValidation(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetPartnerPortOutValidationRequest
		expectedResponse *model.GetPartnerPortOutValidationResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetPartnerPortOutValidationRequest{
				Session:                     createMockSession(),
				GetPartnerPortOutValidation: generateSampleGetPartnerPortOutValidation(),
			},
			expectedResponse: &model.GetPartnerPortOutValidationResponse{
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
					Status      string `xml:"status,attr"`
					Timestamp   string `xml:"timestamp,attr"`
					Type        string `xml:"type,attr"`
					Mdn         string `xml:"mdn"`
					Sim         string `xml:"sim"`
					StatusCode  string `xml:"statusCode"`
					Description string `xml:"description"`
				}{
					Timestamp:   "20200611114406",
					Status:      "success",
					Type:        "QueryPortStatusResponse",
					Mdn:         "545454545454",
					Sim:         "2313254658721",
					StatusCode:  "200",
					Description: "Dummy",
				}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetPartnerPortOutValidationRequest{
				Session:                     createMockSession(),
				GetPartnerPortOutValidation: generateSampleGetPartnerPortOutValidation(),
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
			request: &model.GetPartnerPortOutValidationRequest{
				Session:                     createMockSession(),
				GetPartnerPortOutValidation: generateSampleGetPartnerPortOutValidation(),
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
			request:          &model.GetPartnerPortOutValidationRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetPartnerPortOutValidationRequest.GetPartnerPortOutValidation.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetPartnerPortOutValidationRequest.GetPartnerPortOutValidation.Sim' Error:Field validation for 'Sim' failed on the 'required' tag\nKey: 'GetPartnerPortOutValidationRequest.GetPartnerPortOutValidation.MessageCode' Error:Field validation for 'MessageCode' failed on the 'required' tag\nKey: 'GetPartnerPortOutValidationRequest.GetPartnerPortOutValidation.Description' Error:Field validation for 'Description' failed on the 'required' tag"),
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

			response, err := client.GetPartnerPortOutValidation(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
