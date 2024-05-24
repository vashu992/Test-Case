package model

type GetCancelPortInRequest struct {
	Session          AuthSession `xml:"session"`
	GetMdnSimPayload GetMdnSim   `xml:"request"`
}

type GetCancelPortInResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp   string `xml:"timestamp,attr"`
		Status      string `xml:"status,attr"`
		Type        string `xml:"type,attr"`
		Mdn         string `xml:"mdn" validate:"required,len=10,numeric"`
		Sim         string `xml:"sim" validate:"required,max=25"`
		Result      string `xml:"result"`
		Description string `xml:"description"`
	}
}
