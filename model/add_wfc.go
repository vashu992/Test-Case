package model

type GetAddWFC struct {
	Type        string      `xml:"type,attr"`
	Mdn         string      `xml:"mdn" validate:"required,len=10,numeric"`
	Esn         string      `xml:"esn"`
	E911ADDRESS E911ADDRESS `xml:"e911Address"`
}

type GetAddWFCRequest struct {
	Session   AuthSession `xml:"session"`
	GetAddWFC GetAddWFC   `xml:"request"`
}

type GetAddWFCResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Errors    struct {
			Error struct {
				Code    string `xml:"code"`
				Message string `xml:"message"`
			}
		}
		Type string `xml:"type,attr"`
		Mdn  string `xml:"mdn"`
		Sim  string `xml:"sim"`
	}
}
