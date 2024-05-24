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

func generateSampleGetQueryDevice() model.GetQueryDevice {
	return model.GetQueryDevice{
		Type:  "dummyType",
		IMEI:  "dummyIMEI",
		TAC:   "dummyTAC",
		MODEL: "dummyMODEL",
		NAME:  "dummyNAME",
	}
}

func TestGetQueryDevice(t *testing.T) {

	testCases := []struct {
		name             string
		code             int
		request          *model.GetQueryDeviceRequest
		expectedResponse *model.GetQueryDeviceResponse
		expectedError    error
	}{
		{
			name: validCase,
			request: &model.GetQueryDeviceRequest{
				Session:        createMockSession(),
				GetQueryDevice: generateSampleGetQueryDevice(),
			},
			expectedResponse: &model.GetQueryDeviceResponse{ResponseSession: struct {
				Clec struct {
					ID string `xml:"id"`
				}
				Timestamp string `xml:"timestamp"`
			}{
				Clec: struct {
					ID string `xml:"id"`
				}{ID: "1004"}, Timestamp: "20200611114406"},
				Response: struct {
					Status                string `xml:"status,attr"`
					Timestamp             string `xml:"timestamp,attr"`
					Type                  string `xml:"type,attr"`
					MODEL                 string `xml:"MODEL"`
					IMEI                  string `xml:"IMEI"`
					MARKETINGNAME         string `xml:"MARKETINGNAME"`
					MANUFACTURER          string `xml:"MANUFACTURER"`
					TAC                   string `xml:"TAC"`
					NAME                  string `xml:"NAME"`
					BAND12COMPATIBLE      string `xml:"BAND12COMPATIBLE"`
					VOLTECOMPATIBLE       string `xml:"VOLTECOMPATIBLE"`
					WIFICOMPATIBLE        string `xml:"WIFICOMPATIBLE"`
					NETWORKCOMPATIBLE     string `xml:"NETWORKCOMPATIBLE"`
					NETWORKTECHNOLOGY     string `xml:"NETWORKTECHNOLOGY"`
					TMOBILEAPPROVED       string `xml:"TMOBILEAPPROVED"`
					DUALBANDWIFI          string `xml:"DUALBANDWIFI"`
					IMS                   string `xml:"IMS"`
					IPV6                  string `xml:"IPV6"`
					PASSPOINT             string `xml:"PASSPOINT"`
					ROAMINGIMS            string `xml:"ROAMINGIMS"`
					VOLTEEMERGENCYCALLING string `xml:"VOLTEEMERGENCYCALLING"`
					WIFICALLINGVERSION    string `xml:"WIFICALLINGVERSION"`
					GPRS                  string `xml:"GPRS"`
					HSDPA                 string `xml:"HSDPA"`
					HSPA                  string `xml:"HSPA"`
					VOLTE                 string `xml:"VOLTE"`
					VOWIFI                string `xml:"VOWIFI"`
					WIFI                  string `xml:"WIFI"`
					BANDS                 string `xml:"BANDS"`
					ESIM                  string `xml:"ESIM"`
					REMOTESIMUNLOCK       string `xml:"REMOTESIMUNLOCK"`
					SIMSIZE               string `xml:"SIMSIZE"`
					SIMSLOTS              string `xml:"SIMSLOTS"`
					WLAN                  string `xml:"WLAN"`
					LTE                   string `xml:"LTE"`
					LTEADVANCED           string `xml:"LTEADVANCED"`
					LTECATEGORY           string `xml:"LTECATEGORY"`
					UMTS                  string `xml:"UMTS"`
					CHIPSETMODEL          string `xml:"CHIPSETMODEL"`
					CHIPSETNAME           string `xml:"CHIPSETNAME"`
					CHIPSETVENDOR         string `xml:"CHIPSETVENDOR"`
					DEVICETYPE            string `xml:"DEVICETYPE"`
					OSNAME                string `xml:"OSNAME"`
					PRIMARYHARDWARETYPE   string `xml:"PRIMARYHARDWARETYPE"`
					YEARRELEASED          string `xml:"YEARRELEASED"`
					TECHNOLOGYONTHEDEVICE string `xml:"TECHNOLOGYONTHEDEVICE"`
					HDVOICE               string `xml:"HDVOICE"`
					STANDALONE5G          string `xml:"STANDALONE5G"`
					NONSTANDALONE5G       string `xml:"NONSTANDALONE5G"`
					COMPATIBILITY         string `xml:"COMPATIBILITY"`
					VONRCOMPATIBLE        string `xml:"VONRCOMPATIBLE"`
					StatusCode            string `xml:"statusCode"`
					Description           string `xml:"description"`
				}{
					Status:                "Status",
					Timestamp:             "64654654654",
					Type:                  "active",
					MODEL:                 "123",
					IMEI:                  "5454878564",
					MARKETINGNAME:         "AAA",
					MANUFACTURER:          "ABC",
					TAC:                   "TAC",
					NAME:                  "Name",
					BAND12COMPATIBLE:      "Yes",
					VOLTECOMPATIBLE:       "No",
					WIFICOMPATIBLE:        "No",
					NETWORKCOMPATIBLE:     "5g",
					NETWORKTECHNOLOGY:     "IVR",
					TMOBILEAPPROVED:       "No",
					DUALBANDWIFI:          "No",
					IMS:                   "Active",
					IPV6:                  "No",
					PASSPOINT:             "RRR",
					ROAMINGIMS:            "Yes",
					VOLTEEMERGENCYCALLING: "Yes",
					WIFICALLINGVERSION:    "123",
					GPRS:                  "Yes",
					HSDPA:                 "No",
					HSPA:                  "Yes",
					VOLTE:                 "Yes",
					VOWIFI:                "Yes",
					WIFI:                  "yes",
					BANDS:                 "bands",
					ESIM:                  "5454545454",
					REMOTESIMUNLOCK:       "Yes",
					SIMSIZE:               "34",
					SIMSLOTS:              "2",
					WLAN:                  "13213213",
					LTE:                   "Yes",
					LTEADVANCED:           "Yes",
					LTECATEGORY:           "Yes",
					UMTS:                  "True",
					CHIPSETMODEL:          "123",
					CHIPSETNAME:           "ABC",
					CHIPSETVENDOR:         "XYZ",
					DEVICETYPE:            "ACTIVE",
					OSNAME:                "WIND",
					PRIMARYHARDWARETYPE:   "Phone",
					YEARRELEASED:          "Yes",
					TECHNOLOGYONTHEDEVICE: "ABC",
					HDVOICE:               "Yes",
					STANDALONE5G:          "Yes",
					NONSTANDALONE5G:       "Yes",
					COMPATIBILITY:         "Active",
					VONRCOMPATIBLE:        "No",
					StatusCode:            "200",
					Description:           "Dummy",
				}},
			expectedError: nil,
		},
		{
			name: xmlUnmarshalError,
			request: &model.GetQueryDeviceRequest{
				Session:        createMockSession(),
				GetQueryDevice: generateSampleGetQueryDevice(),
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
			request:          &model.GetQueryDeviceRequest{},
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

			response, err := client.GetQueryDevice(context.TODO(), tc.request)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
