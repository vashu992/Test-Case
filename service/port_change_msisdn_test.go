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

func generateSampleGetPortinWithChangeMSISDN() model.GetPortinWithChangeMSISDN {
	return model.GetPortinWithChangeMSISDN{
		Type: "dummyType",
		PortinWithChangeMSISDNPortInfo: model.PortinWithChangeMSISDNPortInfo{
			Mdn:          "0123456789",
			Newmdn:       "0987654321",
			Esn:          "1234567890123456789012345",
			PlanId:       "plan123",
			Zipcode:      "62701",
			AuthorizedBy: "Jane Doe",
			PortinWithChangeMSISDNBilling: model.PortinWithChangeMSISDNBilling{
				FirstName: "John",
				LastName:  "Doe",
				PortinWithChangeMSISDNAddress: model.PortinWithChangeMSISDNAddress{
					AddressLine1: "123 Main St",
					AddressLine2: "Apt 4B",
					City:         "Springfield",
					State:        "IL",
					Zip:          "62701",
				},
			},
			E911ADDRESS: generateMockE911ADDRESS(),
			PortinWithChangeMSISDNBillingOldProvider: model.PortinWithChangeMSISDNBillingOldProvider{
				Account:  "oldAccount123",
				Password: "oldPassword456",
			},
		},
	}
}

func TestGetPortinWithChangeMSISDN(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetPortinWithChangeMSISDNRequest
		expectedResponse *model.GetPortinWithChangeMSISDNResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetPortinWithChangeMSISDNRequest{
				Session:                   createMockSession(),
				GetPortinWithChangeMSISDN: generateSampleGetPortinWithChangeMSISDN(),
			},
			expectedResponse: &model.GetPortinWithChangeMSISDNResponse{
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
					Timestamp   string `xml:"timestamp,attr"`
					Status      string `xml:"status,attr"`
					Type        string `xml:"type,attr"`
					Mdn         string `xml:"mdn"`
					Sim         string `xml:"sim"`
					Result      string `xml:"result"`
					Description string `xml:"description"`
				}{
					Timestamp:   "20200611114406",
					Status:      "success",
					Type:        "QueryPortStatusResponse",
					Mdn:         "545454545454",
					Sim:         "2313254658721",
					Result:      "Success",
					Description: "Dummy",
				}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetPortinWithChangeMSISDNRequest{
				Session:                   createMockSession(),
				GetPortinWithChangeMSISDN: generateSampleGetPortinWithChangeMSISDN(),
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
			request: &model.GetPortinWithChangeMSISDNRequest{
				Session:                   createMockSession(),
				GetPortinWithChangeMSISDN: generateSampleGetPortinWithChangeMSISDN(),
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
			request:          &model.GetPortinWithChangeMSISDNRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.Newmdn' Error:Field validation for 'Newmdn' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.Esn' Error:Field validation for 'Esn' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.PlanId' Error:Field validation for 'PlanId' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.Zipcode' Error:Field validation for 'Zipcode' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.AuthorizedBy' Error:Field validation for 'AuthorizedBy' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.PortinWithChangeMSISDNBilling.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.PortinWithChangeMSISDNBilling.LastName' Error:Field validation for 'LastName' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.PortinWithChangeMSISDNBilling.PortinWithChangeMSISDNAddress.City' Error:Field validation for 'City' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.PortinWithChangeMSISDNBilling.PortinWithChangeMSISDNAddress.State' Error:Field validation for 'State' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.PortinWithChangeMSISDNBilling.PortinWithChangeMSISDNAddress.Zip' Error:Field validation for 'Zip' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.E911ADDRESS.STREET1' Error:Field validation for 'STREET1' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.E911ADDRESS.CITY' Error:Field validation for 'CITY' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.E911ADDRESS.STATE' Error:Field validation for 'STATE' failed on the 'required' tag\nKey: 'GetPortinWithChangeMSISDNRequest.GetPortinWithChangeMSISDN.PortinWithChangeMSISDNPortInfo.E911ADDRESS.ZIP' Error:Field validation for 'ZIP' failed on the 'required' tag"),
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

			response, err := client.GetPortinWithChangeMSISDN(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
