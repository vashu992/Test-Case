package model

type GetWirelessHotline struct {
	Type              string    `xml:"type,attr"`
	Mdn               string    `xml:"mdn"`
	HotlineNumber     string    `xml:"hotlineNumber"`
	HotlineChargeable string    `xml:"hotlineChargeable"`
	Hotline           []Hotline `xml:"Hotline"`
}

type GetWirelessHotlineRequest struct {
	Session            AuthSession        `xml:"session"`
	GetWirelessHotline GetWirelessHotline `xml:"request"`
}

type GetWirelessHotlineResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		MDN       string `xml:"MDN"`
	}
}
