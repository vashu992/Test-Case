package model

type GetModifyPort struct {
	Type             string           `xml:"type,attr"`
	AgentAccount     string           `xml:"agentAccount"`
	OrderId          string           `xml:"orderId"`
	PortInputChanges PortInputChanges `xml:"portInputChanges"`
}

type PortInputChanges struct {
	Mdn             string          `xml:"mdn" validate:"required,len=10,numeric"`
	Esn             string          `xml:"esn" validate:"required,max=25"`
	AuthorizedBy    string          `xml:"authorizedBy"`
	BusinessName    string          `xml:"businessName"`
	OldProviderPort OldProviderPort `xml:"oldProvider"`
	Billing         Billing         `xml:"billing"`
	E911Address     E911Address     `xml:"e911Address"`
}

type OldProviderPort struct {
	Account  string `xml:"account" validate:"required"`
	Password string `xml:"password" validate:"required"`
	Ssn      string `xml:"ssn"`
	Dob      string `xml:"dob"`
}

type GetModifyPortRequest struct {
	Session       AuthSession   `xml:"session"`
	GetModifyPort GetModifyPort `xml:"request"`
}

type GetModifyPortResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp   string `xml:"timestamp,attr"`
		Status      string `xml:"status,attr"`
		Type        string `xml:"type,attr"`
		Mdn         string `xml:"mdn"`
		Sim         string `xml:"sim"`
		Result      string `xml:"result"`
		Description string `xml:"description"`
		Warning     string `xml:"warning"`
	}
}
