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

func generateSampleGetMdnSim() model.GetMdnSim {
	return model.GetMdnSim{
		Type: "DummyType",
		Mdn:  "1234567890",
		Sim:  "SIM12345678",
	}
}

func TestGetQueryHLR(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetQueryHLRRequest
		expectedResponse *model.GetQueryHLRResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetQueryHLRRequest{
				Session:   createMockSession(),
				GetMdnSim: generateSampleGetMdnSim(),
			},
			expectedResponse: &model.GetQueryHLRResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "212121"},
					Timestamp: "687876546857"},
				Response: struct {
					Status    string `xml:"status,attr"`
					Timestamp string `xml:"timestamp,attr"`
					Type      string `xml:"type,attr"`
					Mdn       string `xml:"mdn"`
					Sim       string `xml:"sim"`
					IMEI      string `xml:"IMEI"`
					IMSI      string `xml:"IMSI"`
					SIMSTATUS string `xml:"SIMSTATUS"`
					Socs      []struct {
						Soc string `xml:"soc"`
					}
					Apn []struct {
						Name  string `xml:"name"`
						Value string `xml:"value"`
					}
					MSSTATUS    string `xml:"MS_STATUS"`
					StatusCode  string `xml:"statusCode"`
					Description string `xml:"description"`
				}{Status: "Active", Timestamp: "544542121545", Type: "LIVE", Mdn: "545421245458", Sim: "7875454578", IMEI: "15454565898665", IMSI: "54545454565655", SIMSTATUS: "ACTIVE",
					Socs: []struct {
						Soc string `xml:"soc"`
					}{
						{Soc: "SOC"},
					},
					Apn: []struct {
						Name  string `xml:"name"`
						Value string `xml:"value"`
					}{
						{
							Name:  "Name",
							Value: "Value",
						},
					}, MSSTATUS: "ACTIVE", StatusCode: "200", Description: "Dummy Description"}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetQueryHLRRequest{
				Session:   createMockSession(),
				GetMdnSim: generateSampleGetMdnSim(),
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
			request:          &model.GetQueryHLRRequest{},
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

			response, err := client.GetQueryHLR(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
