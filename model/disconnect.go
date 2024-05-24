package model

type GetDisconnectRequest struct {
	Session   AuthSession `xml:"session"`
	GetMdnEsn GetMdnEsn   `xml:"request"`
}

type GetDisconnectResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		Mdn       string `xml:"mdn"`
		Esn       string `xml:"esn"`
	}
}
