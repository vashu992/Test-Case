package model

type GetRequestNewSwapMDN struct {
	Type   string `xml:"type,attr"`
	Mdn    string `xml:"mdn"`
	Sim    string `xml:"sim"`
	Imei   string `xml:"imei"`
	NewZip string `xml:"newZip"`
}

type GetRequestNewSwapMDNRequest struct {
	Session              AuthSession          `xml:"session"`
	GetRequestNewSwapMDN GetRequestNewSwapMDN `xml:"request"`
}

type GetRequestNewSwapMDNResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		NewMDN    string `xml:"NewMDN"`
	}
}
