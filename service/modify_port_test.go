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

func generateSampleGetModifyPort() model.GetModifyPort {
	return model.GetModifyPort{
		Type:         "dummy-type",
		AgentAccount: "dummy-agent-account",
		OrderId:      "dummy-order-id",
		PortInputChanges: model.PortInputChanges{
			Mdn:          "1234567890",
			Esn:          "dummy-esn",
			AuthorizedBy: "dummy-authorized-by",
			BusinessName: "dummy-business-name",
			OldProviderPort: model.OldProviderPort{
				Account:  "dummy-account",
				Password: "dummy-password",
				Ssn:      "dummy-ssn",
				Dob:      "dummy-dob",
			},
			Billing:     generateMockBilling(),
			E911Address: generateMockE911Address(),
		},
	}
}

func TestGetModifyPort(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetModifyPortRequest
		expectedResponse *model.GetModifyPortResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetModifyPortRequest{
				Session:       createMockSession(),
				GetModifyPort: generateSampleGetModifyPort(),
			},
			expectedResponse: &model.GetModifyPortResponse{
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
					Warning     string `xml:"warning"`
				}{
					Timestamp:   "20200611114406",
					Status:      "success",
					Type:        "QueryPortStatusResponse",
					Mdn:         "545454545454",
					Sim:         "2313254658721",
					Result:      "Success",
					Description: "Dummy",
					Warning:     "No Warnings",
				},
			},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetModifyPortRequest{
				Session:       createMockSession(),
				GetModifyPort: generateSampleGetModifyPort(),
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
			request:          &model.GetModifyPortRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetModifyPortRequest.GetModifyPort.PortInputChanges.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetModifyPortRequest.GetModifyPort.PortInputChanges.Esn' Error:Field validation for 'Esn' failed on the 'required' tag\nKey: 'GetModifyPortRequest.GetModifyPort.PortInputChanges.OldProviderPort.Account' Error:Field validation for 'Account' failed on the 'required' tag\nKey: 'GetModifyPortRequest.GetModifyPort.PortInputChanges.OldProviderPort.Password' Error:Field validation for 'Password' failed on the 'required' tag\nKey: 'GetModifyPortRequest.GetModifyPort.PortInputChanges.Billing.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag\nKey: 'GetModifyPortRequest.GetModifyPort.PortInputChanges.Billing.LastName' Error:Field validation for 'LastName' failed on the 'required' tag\nKey: 'GetModifyPortRequest.GetModifyPort.PortInputChanges.Billing.Address.Zip' Error:Field validation for 'Zip' failed on the 'required' tag\nKey: 'GetModifyPortRequest.GetModifyPort.PortInputChanges.E911Address.Street1' Error:Field validation for 'Street1' failed on the 'required' tag\nKey: 'GetModifyPortRequest.GetModifyPort.PortInputChanges.E911Address.City' Error:Field validation for 'City' failed on the 'required' tag\nKey: 'GetModifyPortRequest.GetModifyPort.PortInputChanges.E911Address.State' Error:Field validation for 'State' failed on the 'required' tag\nKey: 'GetModifyPortRequest.GetModifyPort.PortInputChanges.E911Address.Zip' Error:Field validation for 'Zip' failed on the 'required' tag"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()

			client := createTestClient(mockServer.URL)

			response, err := client.GetModifyPort(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
