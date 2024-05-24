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

func generateSampleGetValidatePortOutEligibility() model.GetValidatePortOutEligibility {
	return model.GetValidatePortOutEligibility{
		Type:                 "DummyType",
		MDN:                  "1234567890",     // Dummy MDN value
		SIM:                  "SIM1234567890",  // Dummy SIM value
		IMEI:                 "IMEI1234567890", // Dummy IMEI value
		Name:                 "John Doe",       // Dummy name value
		OspAccountNumber:     "1234567",        // Dummy OSP account number
		OspAccountPassword:   "password123",    // Dummy OSP account password
		OspSubscriberAddress: generateMockE911Address(),
	}
}

func TestGetValidatePortOutEligibility(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetValidatePortOutEligibilityRequest
		expectedResponse *model.GetValidatePortOutEligibilityResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetValidatePortOutEligibilityRequest{
				Session:                       createMockSession(),
				GetValidatePortOutEligibility: generateSampleGetValidatePortOutEligibility(),
			},
			expectedResponse: &model.GetValidatePortOutEligibilityResponse{
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
					AttrStatus string `xml:"status,attr"`
					Timestamp  string `xml:"timestamp,attr"`
					Type       string `xml:"type,attr"`
					Result     string `xml:"result"`
					ResultMsg  string `xml:"resultMsg"`
					Status     string `xml:"status"`
				}{AttrStatus: "Available", Timestamp: "20200611114406", Status: "success", Type: "ValidatePortEligibilityResponse", Result: "Done", ResultMsg: "SUCCESS"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetValidatePortOutEligibilityRequest{
				Session:                       createMockSession(),
				GetValidatePortOutEligibility: generateSampleGetValidatePortOutEligibility(),
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
			request: &model.GetValidatePortOutEligibilityRequest{
				Session:                       createMockSession(),
				GetValidatePortOutEligibility: generateSampleGetValidatePortOutEligibility(),
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
			request:          &model.GetValidatePortOutEligibilityRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetValidatePortOutEligibilityRequest.GetValidatePortOutEligibility.MDN' Error:Field validation for 'MDN' failed on the 'required' tag\nKey: 'GetValidatePortOutEligibilityRequest.GetValidatePortOutEligibility.OspAccountNumber' Error:Field validation for 'OspAccountNumber' failed on the 'required' tag\nKey: 'GetValidatePortOutEligibilityRequest.GetValidatePortOutEligibility.OspSubscriberAddress.Street1' Error:Field validation for 'Street1' failed on the 'required' tag\nKey: 'GetValidatePortOutEligibilityRequest.GetValidatePortOutEligibility.OspSubscriberAddress.City' Error:Field validation for 'City' failed on the 'required' tag\nKey: 'GetValidatePortOutEligibilityRequest.GetValidatePortOutEligibility.OspSubscriberAddress.State' Error:Field validation for 'State' failed on the 'required' tag\nKey: 'GetValidatePortOutEligibilityRequest.GetValidatePortOutEligibility.OspSubscriberAddress.Zip' Error:Field validation for 'Zip' failed on the 'required' tag"),
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

			response, err := client.GetValidatePortOutEligibility(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
