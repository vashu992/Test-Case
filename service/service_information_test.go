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

func TestGetServiceInformation(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetServiceInformationRequest
		expectedResponse *model.GetServiceInformationResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetServiceInformationRequest{
				Session:   createMockSession(),
				GetMdnEsn: generateSampleGetMdnEsn(),
			},
			expectedResponse: &model.GetServiceInformationResponse{
				ResponseSession: struct {
					Clec struct {
						ID string `xml:"id"`
					}
					Timestamp string `xml:"timestamp"`
				}{
					Clec: struct {
						ID string `xml:"id"`
					}{
						ID: "1004"}, Timestamp: "20210105080808"},
				Response: struct {
					AttrStatus   string `xml:"status,attr"`
					Timestamp    string `xml:"timestamp,attr"`
					Type         string `xml:"type,attr"`
					Mdn          string `xml:"mdn"`
					Sim          string `xml:"sim"`
					IMEI         string `xml:"IMEI"`
					Status       string `xml:"status"`
					BillCycleDay string `xml:"billCycleDay"`
					Plan         struct {
						Plan          string `xml:"plan"`
						EffectiveDate string `xml:"effectiveDate"`
					}
					Socs []struct {
						Soc           string `xml:"soc"`
						EffectiveDate string `xml:"effectiveDate"`
					}
					BALANCEDETAIL struct {
						SUBSCRIBERSTATE   string `xml:"SUBSCRIBERSTATE"`
						HOTLINENUMBER     string `xml:"HOTLINENUMBER"`
						HOTLINECHARGEABLE string `xml:"HOTLINECHARGEABLE"`
						PURCHASE          []struct {
							PURCHASEID string `xml:"PURCHASEID"`
							TARIFFNAME string `xml:"TARIFFNAME"`
							PLANCODE   string `xml:"PLANCODE"`
							BALANCES   struct {
								SUBSCRIPTIONID string `xml:"SUBSCRIPTIONID"`
								UOM            string `xml:"UOM"`
								BALANCE        string `xml:"BALANCE"`
								VALIDFROM      string `xml:"VALIDFROM"`
								VALIDTO        string `xml:"VALIDTO"`
							}
						}
						TOTALBALANCE struct {
							TALK string `xml:"TALK"`
							TEXT string `xml:"TEXT"`
							DATA string `xml:"DATA"`
						}
					}
					StatusCode  string `xml:"statusCode"`
					Description string `xml:"description"`
				}{
					AttrStatus:   "Active",
					Timestamp:    "5454878787",
					Type:         "ABC",
					Mdn:          "554545451215",
					Sim:          "58415245215",
					IMEI:         "1212554656569",
					Status:       "Active",
					BillCycleDay: "20",
					Plan: struct {
						Plan          string `xml:"plan"`
						EffectiveDate string `xml:"effectiveDate"`
					}{Plan: "Super", EffectiveDate: "20"},
					Socs: []struct {
						Soc           string `xml:"soc"`
						EffectiveDate string `xml:"effectiveDate"`
					}{
						{
							Soc:           "SOC",
							EffectiveDate: "20",
						},
					},
					BALANCEDETAIL: struct {
						SUBSCRIBERSTATE   string `xml:"SUBSCRIBERSTATE"`
						HOTLINENUMBER     string `xml:"HOTLINENUMBER"`
						HOTLINECHARGEABLE string `xml:"HOTLINECHARGEABLE"`
						PURCHASE          []struct {
							PURCHASEID string `xml:"PURCHASEID"`
							TARIFFNAME string `xml:"TARIFFNAME"`
							PLANCODE   string `xml:"PLANCODE"`
							BALANCES   struct {
								SUBSCRIPTIONID string `xml:"SUBSCRIPTIONID"`
								UOM            string `xml:"UOM"`
								BALANCE        string `xml:"BALANCE"`
								VALIDFROM      string `xml:"VALIDFROM"`
								VALIDTO        string `xml:"VALIDTO"`
							}
						}
						TOTALBALANCE struct {
							TALK string `xml:"TALK"`
							TEXT string `xml:"TEXT"`
							DATA string `xml:"DATA"`
						}
					}{SUBSCRIBERSTATE: "ACTIVE", HOTLINENUMBER: "123", HOTLINECHARGEABLE: "Yes", PURCHASE: []struct {
						PURCHASEID string `xml:"PURCHASEID"`
						TARIFFNAME string `xml:"TARIFFNAME"`
						PLANCODE   string `xml:"PLANCODE"`
						BALANCES   struct {
							SUBSCRIPTIONID string `xml:"SUBSCRIPTIONID"`
							UOM            string `xml:"UOM"`
							BALANCE        string `xml:"BALANCE"`
							VALIDFROM      string `xml:"VALIDFROM"`
							VALIDTO        string `xml:"VALIDTO"`
						}
					}{
						{
							PURCHASEID: "124564",
							TARIFFNAME: "SUPEER",
							PLANCODE:   "011",
							BALANCES: struct {
								SUBSCRIPTIONID string `xml:"SUBSCRIPTIONID"`
								UOM            string `xml:"UOM"`
								BALANCE        string `xml:"BALANCE"`
								VALIDFROM      string `xml:"VALIDFROM"`
								VALIDTO        string `xml:"VALIDTO"`
							}{
								SUBSCRIPTIONID: "5445451215",
								UOM:            "12121",
								BALANCE:        "54",
								VALIDFROM:      "05/05/2024",
								VALIDTO:        "05/05/2025",
							},
						},
					}, TOTALBALANCE: struct {
						TALK string `xml:"TALK"`
						TEXT string `xml:"TEXT"`
						DATA string `xml:"DATA"`
					}{TALK: "50MIN", TEXT: "50", DATA: "1024"}},
					StatusCode:  "200",
					Description: "Dummy",
				}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetServiceInformationRequest{
				Session:   createMockSession(),
				GetMdnEsn: generateSampleGetMdnEsn(),
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
			request:          &model.GetServiceInformationRequest{},
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

			response, err := client.GetServiceInformation(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
