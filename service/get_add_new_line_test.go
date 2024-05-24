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

func generateSampleGetAddNewLine() model.GetAddNewLine {
	return model.GetAddNewLine{
		Type: "dummy-type",
		AddNewLineShared: model.AddNewLineShared{
			ParentMDN: "1234567890",
			ParentSIM: "dummy-parent-SIM",
			LineDetails: model.LineDetails{
				Line: model.Line{
					MDN: "1234567890",
					SIM: "dummy-SIM",
				},
			},
		},
	}
}

func TestGetAddNewLine(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetAddNewLineRequest
		expectedResponse *model.GetAddNewLineResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetAddNewLineRequest{

				Session:       createMockSession(),
				GetAddNewLine: generateSampleGetAddNewLine(),
			},
			expectedResponse: &model.GetAddNewLineResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20210405142409"},
				Response: struct {
					Timestamp string `xml:"timestamp,attr"`
					Status    string `xml:"status,attr"`
					Type      string `xml:"type,attr"`
					MDN       string `xml:"MDN"`
				}{Timestamp: "20210405142409", Status: "success", Type: "AddNewLineResponse", MDN: "1234567890"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetAddNewLineRequest{

				Session:       createMockSession(),
				GetAddNewLine: generateSampleGetAddNewLine(),
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
			request: &model.GetAddNewLineRequest{

				Session:       createMockSession(),
				GetAddNewLine: generateSampleGetAddNewLine(),
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
			request:          &model.GetAddNewLineRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetAddNewLineRequest.GetAddNewLine.AddNewLineShared.ParentMDN' Error:Field validation for 'ParentMDN' failed on the 'required' tag\nKey: 'GetAddNewLineRequest.GetAddNewLine.AddNewLineShared.LineDetails.Line.MDN' Error:Field validation for 'MDN' failed on the 'required' tag"),
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

			response, err := client.GetAddNewLine(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
