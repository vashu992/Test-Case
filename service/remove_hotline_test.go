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

func generateSampleGetRemoveHotline() model.GetRemoveHotline {
	return model.GetRemoveHotline{
		Text: "SomeText",
		Type: "DummyType",
		Mdn:  "1234567890", // Dummy MDN value
		Hotline: model.Hotline{
			Service: "DummyService", // Dummy service value
		},
	}
}

func TestGetRemoveHotline(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetRemoveHotlineRequest
		expectedResponse *model.GetRemoveHotlineResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetRemoveHotlineRequest{
				Session:          createMockSession(),
				GetRemoveHotline: generateSampleGetRemoveHotline(),
			},
			expectedResponse: &model.GetRemoveHotlineResponse{
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
					Timestamp string `xml:"timestamp,attr"`
					Status    string `xml:"status,attr"`
					Type      string `xml:"type,attr"`
					MDN       string `xml:"MDN"`
				}{Timestamp: "20200611114406", Status: "success", Type: "RemoveHotLineResponse", MDN: "1234567890"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetRemoveHotlineRequest{
				Session:          createMockSession(),
				GetRemoveHotline: generateSampleGetRemoveHotline(),
			},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
				Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
				Err:     errors.New("EOF"),
			},
		},
		{
			name:             internalServerError,
			request:          &model.GetRemoveHotlineRequest{},
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
			response, err := client.GetRemoveHotline(context.TODO(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
