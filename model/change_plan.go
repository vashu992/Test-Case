package model

type GetChangePlan struct {
	Type        string      `xml:"type,attr"`
	Mdn         string      `xml:"mdn" validate:"required,len=10"`
	Sim         string      `xml:"sim" validate:"alphanum,max=25"`
	NewplanID   string      `xml:"newplanID" validate:"required"`
	E911ADDRESS E911ADDRESS `xml:"E911ADDRESS"`
}

type GetChangePlanRequest struct {
	Session       AuthSession   `xml:"session"`
	GetChangePlan GetChangePlan `xml:"request"`
}

type GetChangePlanResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		Message   string `xml:"message"`
		Warning   string `xml:"warning"`
	}
}
