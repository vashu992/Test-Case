package model

type GetCancelDeviceLocation struct {
	Type string `xml:"type,attr"`
	Mdn  string `xml:"mdn" validate:"required,len=10,numeric"`
	Esn  string `xml:"esn"`
}

type GetCancelDeviceLocationRequest struct {
	Session                 AuthSession             `xml:"session"`
	GetCancelDeviceLocation GetCancelDeviceLocation `xml:"request"`
}

type GetCancelDeviceLocationResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		Message   string `xml:"message"`
		MDN       string `xml:"MDN"`
		ESN       string `xml:"ESN"`
	}
}
