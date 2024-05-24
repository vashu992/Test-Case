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

func generateSampleGetChangeVoicemailLanguage() model.GetChangeVoicemailLanguage {
	return model.GetChangeVoicemailLanguage{
		Type:     "dummy-type",
		Mdn:      "1234567890",
		Sim:      "dummy-sim",
		Language: "EN",
	}
}

func TestGetChangeVoicemailLanguage(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetChangeVoicemailLanguageRequest
		expectedResponse *model.GetChangeVoicemailLanguageResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetChangeVoicemailLanguageRequest{
				Session:                    createMockSession(),
				GetChangeVoicemailLanguage: generateSampleGetChangeVoicemailLanguage(),
			},
			expectedResponse: &model.GetChangeVoicemailLanguageResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20210111034011"},
				Response: struct {
					Timestamp string `xml:"timestamp,attr"`
					Status    string `xml:"status,attr"`
					Type      string `xml:"type,attr"`
					Mdn       string `xml:"mdn"`
					Sim       string `xml:"sim"`
				}{Timestamp: "20210111034011", Status: "success", Type: "ResetVoicemailPasswordResponse", Mdn: "12345534545", Sim: "2343244322343243234"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetChangeVoicemailLanguageRequest{
				Session:                    createMockSession(),
				GetChangeVoicemailLanguage: generateSampleGetChangeVoicemailLanguage(),
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
			request: &model.GetChangeVoicemailLanguageRequest{
				Session:                    createMockSession(),
				GetChangeVoicemailLanguage: generateSampleGetChangeVoicemailLanguage(),
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
			request:          &model.GetChangeVoicemailLanguageRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetChangeVoicemailLanguageRequest.GetChangeVoicemailLanguage.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetChangeVoicemailLanguageRequest.GetChangeVoicemailLanguage.Language' Error:Field validation for 'Language' failed on the 'required' tag"),
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

			response, err := client.GetChangeVoicemailLanguage(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
