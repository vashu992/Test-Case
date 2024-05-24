package model

type GetRequestNewSwapESN struct {
	Type   string `xml:"type,attr"`
	Mdn    string `xml:"mdn"`
	Newesn string `xml:"newesn"`
	Oldesn string `xml:"oldesn"`
}

type GetRequestNewSwapESNRequest struct {
	Session              AuthSession          `xml:"session"`
	GetRequestNewSwapESN GetRequestNewSwapESN `xml:"request"`
}

type GetRequestNewSwapESNResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		Mdn       string `xml:"mdn"`
		Esn       string `xml:"esn"`
	}
}
