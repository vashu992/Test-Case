package model

type GetRemoveHotline struct {
	Text    string  `xml:",chardata"`
	Type    string  `xml:"type,attr"`
	Mdn     string  `xml:"mdn"`
	Hotline Hotline `xml:"Hotline"`
}

type GetRemoveHotlineRequest struct {
	Session          AuthSession      `xml:"session"`
	GetRemoveHotline GetRemoveHotline `xml:"request"`
}

type GetRemoveHotlineResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		MDN       string `xml:"MDN"`
	}
}
