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

func generateSampleGetESIMProfile() model.GetESIMProfile {
	return model.GetESIMProfile{
		Sim:       "dummy-sim",
		ESimBrand: "dummy-eSimBrand",
	}
}

func TestGetESIMProfile(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetESIMProfileRequest
		expectedResponse *model.GetESIMProfileResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetESIMProfileRequest{

				Session:        createMockSession(),
				GetESIMProfile: generateSampleGetESIMProfile(),
			},
			expectedResponse: &model.GetESIMProfileResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20230817151512"},
				Response: struct {
					Status         string `xml:"status,attr"`
					Timestamp      string `xml:"timestamp,attr"`
					Type           string `xml:"type,attr"`
					StatusCode     string `xml:"statusCode"`
					Description    string `xml:"description"`
					Iccid          string `xml:"iccid"`
					ActivationCode string `xml:"activationCode"`
					ProfileType    string `xml:"profileType"`
					Lastmodified   string `xml:"lastmodified"`
					MatchingId     string `xml:"matchingId"`
					ProfileState   string `xml:"profileState"`
					ESimBrand      string `xml:"eSimBrand"`
				}{Status: "success", Timestamp: "20230817151512", Type: "", StatusCode: "00", Description: "SUCCESS", Iccid: "8901240347177226293", ActivationCode: "1$t-mobile.idemia.io$Y3ALS-UEIZ3-8FAPD-FZY72", ProfileType: "CONV5GPWGQRID", Lastmodified: "2023-06-19T10:03:17Z", MatchingId: "Y3ALS-UEIZ3-8FAPD-FZY72", ProfileState: "RELEASED", ESimBrand: "PWGQR"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetESIMProfileRequest{

				Session:        createMockSession(),
				GetESIMProfile: generateSampleGetESIMProfile(),
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
			request: &model.GetESIMProfileRequest{
				Session:        createMockSession(),
				GetESIMProfile: generateSampleGetESIMProfile(),
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
			request:          &model.GetESIMProfileRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetESIMProfileRequest.GetESIMProfile.ESimBrand' Error:Field validation for 'ESimBrand' failed on the 'required' tag"),
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

			response, err := client.GetESIMProfile(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
