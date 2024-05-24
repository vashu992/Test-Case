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

func generateSampleGetCoverage2() model.GetCoverage2 {
	return model.GetCoverage2{
		Type: "dummy-type",
		Zip:  "12345",
	}
}

func TestGetCoverage2(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetCoverage2Request
		expectedResponse *model.GetCoverage2Response
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetCoverage2Request{

				Session:      createMockSession(),
				GetCoverage2: generateSampleGetCoverage2(),
			},
			expectedResponse: &model.GetCoverage2Response{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20150608170550"},
				Response: struct {
					Status             string `xml:"status,attr"`
					Timestamp          string `xml:"timestamp,attr"`
					Type               string `xml:"type,attr"`
					Zip                string `xml:"zip"`
					StatusCode         string `xml:"statusCode"`
					ActivationEligible string `xml:"activationEligible"`
				}{Status: "success", Timestamp: "20150608170550", Type: "GetCoverageResponse", Zip: "83501", StatusCode: "00", ActivationEligible: "True"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetCoverage2Request{

				Session:      createMockSession(),
				GetCoverage2: generateSampleGetCoverage2(),
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
			request: &model.GetCoverage2Request{
				Session:      createMockSession(),
				GetCoverage2: generateSampleGetCoverage2(),
			},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_FAILED_CODE,
				Message: constants.REQUEST_FAILED_INFO,
				Err:     errors.New("Post \"\": unsupported protocol scheme \"\""),
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
			response, err := client.GetCoverage2(context.TODO(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				return
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
