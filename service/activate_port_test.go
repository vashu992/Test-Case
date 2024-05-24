package pwg_test

import (
	"context"
	"errors"
	"net/http"
	constants "propwg/common"
	util "propwg/common"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vashu992/Test-Case/model"
)

func generateMockE911Address() model.E911Address {
	return model.E911Address{
		Street1: "some streat1",
		Street2: "some streat 2",
		City:    "city",
		State:   "state",
		Zip:     "12345",
	}
}

func generateMockE911ADDRESS() model.E911ADDRESS {
	return model.E911ADDRESS{
		STREET1: "some streat1",
		STREET2: "some streat 2",
		CITY:    "city",
		STATE:   "state",
		ZIP:     "12345",
	}
}

func generateMockBilling() model.Billing {
	return model.Billing{
		FirstName: "first name",
		LastName:  "last name",
		Address: model.Address{
			Zip: "12345",
		},
	}
}

func generateMockActivatePortIn() model.ActivatePortIn {
	return model.ActivatePortIn{
		// Type:         "dummy-type",
		ActivityType: "dummy-activity-type",
		Esn:          "dummy-esn",
		Ssn:          "dummy-ssn",
		Dob:          "dummy-dob",
		PlanId:       "dummy-plan-id",
		BillingCode:  "dummy-billing-code",
		PortInfo: model.PortInfo{
			Mdn:          "1234567890",
			AuthorizedBy: "dummy-authorized-by",
			Billing:      generateMockBilling(),
			OldProvider: model.OldProvider{
				Account:  "dummy-account",
				Password: "dummy-password",
				Esn:      "dummy-old-provider-esn",
			},
			E911Address: generateMockE911Address(),
		},
	}
}

func TestActivatePortIn(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.ActivatePortInRequest
		expectedResponse *model.ActivatePortInResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.ActivatePortInRequest{
				Session:        createMockSession(),
				ActivatePortIn: generateMockActivatePortIn(),
			},
			code: http.StatusOK,
			expectedResponse: &model.ActivatePortInResponse{
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
				}{Timestamp: "20200611114406", Status: "success", Type: "WirelessActivateResponse", Mdn: "1234567890", Sim: "89012607230000000", Result: "", Description: "SUCCESS", Warning: "some warning"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.ActivatePortInRequest{
				Session:        createMockSession(),
				ActivatePortIn: generateMockActivatePortIn(),
			},
			code:             http.StatusOK,
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
				Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
				Err:     errors.New("EOF"),
			},
		},
		{
			name:             inValidCase,
			request:          &model.ActivatePortInRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'ActivatePortInRequest.ActivatePortIn.ActivityType' Error:Field validation for 'ActivityType' failed on the 'required' tag\nKey: 'ActivatePortInRequest.ActivatePortIn.Esn' Error:Field validation for 'Esn' failed on the 'required' tag\nKey: 'ActivatePortInRequest.ActivatePortIn.PlanId' Error:Field validation for 'PlanId' failed on the 'required' tag\nKey: 'ActivatePortInRequest.ActivatePortIn.PortInfo.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'ActivatePortInRequest.ActivatePortIn.PortInfo.Billing.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag\nKey: 'ActivatePortInRequest.ActivatePortIn.PortInfo.Billing.LastName' Error:Field validation for 'LastName' failed on the 'required' tag\nKey: 'ActivatePortInRequest.ActivatePortIn.PortInfo.Billing.Address.Zip' Error:Field validation for 'Zip' failed on the 'required' tag\nKey: 'ActivatePortInRequest.ActivatePortIn.PortInfo.E911Address.Street1' Error:Field validation for 'Street1' failed on the 'required' tag\nKey: 'ActivatePortInRequest.ActivatePortIn.PortInfo.E911Address.City' Error:Field validation for 'City' failed on the 'required' tag\nKey: 'ActivatePortInRequest.ActivatePortIn.PortInfo.E911Address.State' Error:Field validation for 'State' failed on the 'required' tag\nKey: 'ActivatePortInRequest.ActivatePortIn.PortInfo.E911Address.Zip' Error:Field validation for 'Zip' failed on the 'required' tag"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()
			client := createTestClient(mockServer.URL)
			response, err := client.ActivatePortIn(context.Background(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
