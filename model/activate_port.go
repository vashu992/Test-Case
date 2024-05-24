package model

type OldProvider struct {
	Account  string `xml:"account"`
	Password string `xml:"password"`
	Esn      string `xml:"esn" validate:"max=25"`
}

type PortInfo struct {
	Mdn          string      `xml:"mdn" validate:"required,len=10"`
	AuthorizedBy string      `xml:"authorizedBy"`
	Billing      Billing     `xml:"billing"`
	E911Address  E911Address `xml:"E911ADDRESS"`
	OldProvider  OldProvider `xml:"oldProvider"`
}

type ActivatePortIn struct {
	ActivityType string   `xml:"activityType" validate:"required"`
	Esn          string   `xml:"esn" validate:"required,max=25"`
	Ssn          string   `xml:"ssn" validate:"max=9"`
	Dob          string   `xml:"dob"`
	PlanId       string   `xml:"planId" validate:"required"`
	BillingCode  string   `xml:"BillingCode"`
	PortInfo     PortInfo `xml:"portInfo"`
}

type ActivatePortInRequest struct {
	Session        AuthSession    `xml:"session"`
	ActivatePortIn ActivatePortIn `xml:"request"`
}

type ActivatePortInResponse struct {
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
