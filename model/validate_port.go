package model

type GetValidatePortRequest struct {
	Session   AuthSession `xml:"session"`
	GetMdnSim GetMdnSim   `xml:"request"`
}

type GetValidatePortResponse struct {
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
