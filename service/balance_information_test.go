package pwg_test

import (
	"context"
	"errors"
	"net/http"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateSampleGetBalanceInformation() model.GetBalanceInformation {
	return model.GetBalanceInformation{
		Type:           "dummy-type",
		Mdn:            "1234567890",              // Validating for 10 characters
		Esn:            "dummyESN123456789012345", // Validating for alphanumeric and max 25 characters
		PendingBalance: "100.00",
	}
}

func TestGetBalanceInformation(t *testing.T) {
	testCases := []struct {
		name             string
		code             int
		request          *model.GetBalanceInformationRequest
		expectedResponse *model.GetBalanceInformationResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetBalanceInformationRequest{
				Session:               createMockSession(),
				GetBalanceInformation: generateSampleGetBalanceInformation(),
			},
			expectedResponse: &model.GetBalanceInformationResponse{
				ResponseSession: model.ResponseSession{},
				Response: struct {
					AttrStatus    string "xml:\"status,attr\""
					Timestamp     string "xml:\"timestamp,attr\""
					Type          string "xml:\"type,attr\""
					Mdn           string "xml:\"mdn\""
					Sim           string "xml:\"sim\""
					IMEI          string "xml:\"IMEI\""
					Status        string "xml:\"status\""
					Customerid    string "xml:\"customerid\""
					OUTOFCREDIT   string "xml:\"OUTOFCREDIT\""
					BillCycleDay  string "xml:\"billCycleDay\""
					BALANCEDETAIL struct {
						HOTLINENUMBER     string "xml:\"HOTLINENUMBER\""
						HOTLINECHARGEABLE string "xml:\"HOTLINECHARGEABLE\""
						PURCHASE          struct {
							PURCHASEID string "xml:\"PURCHASEID\""
							TARIFFNAME string "xml:\"TARIFFNAME\""
							PLANCODE   string "xml:\"PLANCODE\""
							BALANCES   []struct {
								SUBSCRIPTIONID string "xml:\"SUBSCRIPTIONID\""
								UOM            string "xml:\"UOM\""
								BALANCE        string "xml:\"BALANCE\""
								VALIDFROM      string "xml:\"VALIDFROM\""
								VALIDTO        string "xml:\"VALIDTO\""
							}
						}
						TOTALBALANCE struct {
							TALK string "xml:\"TALK\""
							TEXT string "xml:\"TEXT\""
							DATA string "xml:\"DATA\""
						}
					}
					StatusCode  string "xml:\"statusCode\""
					Description string "xml:\"description\""
				}{
					AttrStatus:   "",
					Timestamp:    "20150504131911",
					Type:         "WirelessAdjustBalanceResponse",
					Mdn:          "4257727633",
					Sim:          "890126070000000111111",
					IMEI:         "",
					Status:       "success",
					Customerid:   "xxxxx",
					OUTOFCREDIT:  "CALLS BEING REDIRECTED TO OUT OF CREDIT NUMBER",
					BillCycleDay: "19",
					StatusCode:   "SUCCESS",
					Description:  "Successful Subscriber Inquiry",
				},
			},
			expectedError: nil,
			code:          http.StatusOK,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetBalanceInformationRequest{
				Session:               createMockSession(),
				GetBalanceInformation: generateSampleGetBalanceInformation(),
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
			code:             http.StatusInternalServerError,
			request:          &model.GetBalanceInformationRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetBalanceInformationRequest.GetBalanceInformation.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetBalanceInformationRequest.GetBalanceInformation.Esn' Error:Field validation for 'Esn' failed on the 'alphanum' tag"),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()
			client := createTestClient(mockServer.URL)
			response, err := client.GetBalanceInformation(context.TODO(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}

func TestGetBalanceInformationWithoutCredit(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetBalanceInformationRequest
		expectedResponse *model.GetBalanceInformationWithoutCreditResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetBalanceInformationRequest{
				Session:               createMockSession(),
				GetBalanceInformation: generateSampleGetBalanceInformation(),
			},
			expectedResponse: &model.GetBalanceInformationWithoutCreditResponse{},
			expectedError:    nil,
			code:             http.StatusOK,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetBalanceInformationRequest{
				Session:               createMockSession(),
				GetBalanceInformation: generateSampleGetBalanceInformation(),
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
			code:             http.StatusInternalServerError,
			request:          &model.GetBalanceInformationRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetBalanceInformationRequest.GetBalanceInformation.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetBalanceInformationRequest.GetBalanceInformation.Esn' Error:Field validation for 'Esn' failed on the 'alphanum' tag"),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()
			client := createTestClient(mockServer.URL)
			response, err := client.GetBalanceInformationWithoutCredit(context.TODO(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}

func TestGetBalanceInformationPendingBalance(t *testing.T) {
	testCases := []struct {
		name             string
		code             int
		request          *model.GetBalanceInformationRequest
		expectedResponse *model.GetBalanceInformationPendingBalanceResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetBalanceInformationRequest{
				Session:               createMockSession(),
				GetBalanceInformation: generateSampleGetBalanceInformation(),
			},
			expectedResponse: &model.GetBalanceInformationPendingBalanceResponse{},
			expectedError:    nil,
			code:             http.StatusOK,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetBalanceInformationRequest{
				Session:               createMockSession(),
				GetBalanceInformation: generateSampleGetBalanceInformation(),
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
			code:             http.StatusInternalServerError,
			request:          &model.GetBalanceInformationRequest{},
			expectedResponse: nil,
			expectedError: &util.WrappedError{
				Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
				Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
				Err:     errors.New("Key: 'GetBalanceInformationRequest.GetBalanceInformation.Mdn' Error:Field validation for 'Mdn' failed on the 'required' tag\nKey: 'GetBalanceInformationRequest.GetBalanceInformation.Esn' Error:Field validation for 'Esn' failed on the 'alphanum' tag"),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockServer := createMockServer(tc.expectedResponse, tc.code)
			defer mockServer.Close()
			client := createTestClient(mockServer.URL)
			response, err := client.GetBalanceInformationPendingBalance(context.TODO(), tc.request)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
