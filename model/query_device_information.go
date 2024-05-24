package model

type GetQueryDevice struct {
	Type  string `xml:"type,attr"`
	IMEI  string `xml:"IMEI"`
	TAC   string `xml:"TAC"`
	MODEL string `xml:"MODEL"`
	NAME  string `xml:"NAME"`
}

type GetQueryDeviceRequest struct {
	Session        AuthSession    `xml:"session"`
	GetQueryDevice GetQueryDevice `xml:"request"`
}

type GetQueryDeviceResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
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
	}
}
