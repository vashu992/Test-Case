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

func generateSampleCoverage() model.Coverage {
	return model.Coverage{
		Carrier: "dummy-carrier",
		Zip:     "1234567",
		Type:    "dummy-type",
	}
}

func TestGetCoverage(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetCoverageRequest
		expectedResponse *model.GetCoverageResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetCoverageRequest{
				CoverageRequest: model.CoverageRequest{
					Session: createMockSession(),
					Request: generateSampleCoverage(),
				},
			},
			expectedResponse: &model.GetCoverageResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20210105080808"},
				Response: struct {
					Status              string   `xml:"status,attr"`
					Timestamp           string   `xml:"timestamp,attr"`
					Type                string   `xml:"type,attr"`
					Zip                 string   `xml:"zip"`
					StatusCode          string   `xml:"statusCode"`
					CoverageQualityIden []string `xml:"coverageQualityIden"`
					Csa                 []string `xml:"csa"`
				}{Status: "success", Timestamp: "20210105080808", Type: "GetCoverageResponse2", Zip: "98104", StatusCode: "00", CoverageQualityIden: []string(nil), Csa: []string(nil)}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetCoverageRequest{
				CoverageRequest: model.CoverageRequest{
					Session: createMockSession(),
					Request: generateSampleCoverage(),
				},
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
			request:          &model.GetCoverageRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetCoverageRequest.CoverageRequest.Request.Carrier' Error:Field validation for 'Carrier' failed on the 'required' tag\nKey: 'GetCoverageRequest.CoverageRequest.Request.Zip' Error:Field validation for 'Zip' failed on the 'required' tag"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()

			client := createTestClient(mockServer.URL)

			response, err := client.GetCoverage(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
