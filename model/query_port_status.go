package model

type GetQueryPortInStatus struct { // in common
	Type string `xml:"type,attr"`
	Mdn  string `xml:"mdn"`
}

type GetQueryPortInStatusRequest struct {
	Session              AuthSession          `xml:"session"`
	GetQueryPortInStatus GetQueryPortInStatus `xml:"request"`
}

type GetQueryPortInStatusResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp   string `xml:"timestamp,attr"`
		Status      string `xml:"status,attr"`
		Type        string `xml:"type,attr"`
		Mdn         string `xml:"mdn"`
		Result      string `xml:"result"`
		Description string `xml:"description"`
	}
}
