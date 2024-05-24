package model

type GetUpdateAddress struct {
	Type        string      `xml:"type,attr"`
	Mdn         string      `xml:"mdn" validate:"required,len=10,numeric"`
	Esn         string      `xml:"esn"`
	E911Address E911Address `xml:"e911Address"`
}

type GetUpdateAddressRequest struct {
	Session          AuthSession      `xml:"session"`
	GetUpdateAddress GetUpdateAddress `xml:"request"`
}

type GetUpdateAddressResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Errors    struct {
			Text  string `xml:",chardata"`
			Error struct {
				Code    string `xml:"code"`
				Message string `xml:"message"`
			}
		}
		Mdn string `xml:"mdn"`
		Sim string `xml:"sim"`
	}
}
