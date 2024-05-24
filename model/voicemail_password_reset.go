package model

type GetVoicemailPasswordResetRequest struct {
	Session   AuthSession `xml:"session"`
	GetMdnSim GetMdnSim   `xml:"request"`
}

type GetVoicemailPasswordResetResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		Mdn       string `xml:"mdn"`
		Sim       string `xml:"sim"`
	}
}
