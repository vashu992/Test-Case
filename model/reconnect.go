package model

type GetReconnect struct {
	Type        string      `xml:"type,attr"`
	Mdn         string      `xml:"mdn"`
	Esn         string      `xml:"esn"`
	Imei        string      `xml:"imei"`
	Plan        string      `xml:"plan"`
	Zip         string      `xml:"zip"`
	BillingCode string      `xml:"BillingCode"`
	StateCode   string      `xml:"stateCode"`
	E911ADDRESS E911ADDRESS `xml:"E911ADDRESS"`
}

type GetReconnectRequest struct {
	Session      AuthSession  `xml:"session"`
	StateCode    string       `xml:"stateCode"`
	GetReconnect GetReconnect `xml:"request"`
}

type GetReconnectResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		Mdn       string `xml:"mdn"`
		Esn       string `xml:"esn"`
		Warning   string `xml:"warning"`
	}
}
