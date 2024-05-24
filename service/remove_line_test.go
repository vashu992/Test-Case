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

func generateSampleGetRemoveLine() model.GetRemoveLine {
	return model.GetRemoveLine{
		Type: "DummyType",
		AddNewLineShared: model.AddNewLineShared{
			ParentMDN: "1234567890",    // Dummy Parent MDN value
			ParentSIM: "SIM1234567890", // Dummy Parent SIM value
			LineDetails: model.LineDetails{
				Line: model.Line{
					MDN: "0987654321",    // Dummy Line MDN value
					SIM: "SIM0987654321", // Dummy Line SIM value
				},
			},
		},
	}
}

func TestGetRemoveLine(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetRemoveLineRequest
		expectedResponse *model.GetRemoveLineResponse
		expectedError    error
	}{
		{
			name: "Valid request",
			request: &model.GetRemoveLineRequest{
				Session:       createMockSession(),
				GetRemoveLine: generateSampleGetRemoveLine(),
			},
			expectedResponse: &model.GetRemoveLineResponse{
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
				}{Timestamp: "20200611114406", Status: "success", Type: "RemoveLineTestResponse", MDN: "1234567890"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetRemoveLineRequest{
				Session:       createMockSession(),
				GetRemoveLine: generateSampleGetRemoveLine(),
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
			request: &model.GetRemoveLineRequest{
				Session:       createMockSession(),
				GetRemoveLine: generateSampleGetRemoveLine(),
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
			request:          &model.GetRemoveLineRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetRemoveLineRequest.GetRemoveLine.AddNewLineShared.ParentMDN' Error:Field validation for 'ParentMDN' failed on the 'required' tag\nKey: 'GetRemoveLineRequest.GetRemoveLine.AddNewLineShared.LineDetails.Line.MDN' Error:Field validation for 'MDN' failed on the 'required' tag"),
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

			response, err := client.GetRemoveLine(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
