package model

type GetPortinWithChangeMSISDN struct {
	Type                           string                         `xml:"type,attr"`
	PortinWithChangeMSISDNPortInfo PortinWithChangeMSISDNPortInfo `xml:"portInfo"`
}

type PortinWithChangeMSISDNAddress struct {
	AddressLine1 string `xml:"addressLine1"`
	AddressLine2 string `xml:"addressLine2"`
	City         string `xml:"city" validate:"required"`
	State        string `xml:"state" validate:"required"`
	Zip          string `xml:"zip" validate:"required,len=5,numeric"`
}

type PortinWithChangeMSISDNPortInfo struct {
	Mdn                                      string                                   `xml:"mdn" validate:"required,len=10,numeric"`
	Newmdn                                   string                                   `xml:"newmdn" validate:"required,len=10,numeric"`
	Esn                                      string                                   `xml:"esn" validate:"required,max=25"`
	PlanId                                   string                                   `xml:"planId" validate:"required"`
	Zipcode                                  string                                   `xml:"zipcode" validate:"required,len=5,numeric"`
	AuthorizedBy                             string                                   `xml:"authorizedBy" validate:"required"`
	PortinWithChangeMSISDNBilling            PortinWithChangeMSISDNBilling            `xml:"billing"`
	E911ADDRESS                              E911ADDRESS                              `xml:"E911ADDRESS"`
	PortinWithChangeMSISDNBillingOldProvider PortinWithChangeMSISDNBillingOldProvider `xml:"oldProvider"`
}

type PortinWithChangeMSISDNBillingOldProvider struct {
	Account  string `xml:"account"`
	Password string `xml:"password"`
}
type PortinWithChangeMSISDNBilling struct {
	FirstName                     string                        `xml:"firstName" validate:"required"`
	LastName                      string                        `xml:"lastName" validate:"required"`
	PortinWithChangeMSISDNAddress PortinWithChangeMSISDNAddress `xml:"address"`
}

type GetPortinWithChangeMSISDNRequest struct {
	Session                   AuthSession               `xml:"session"`
	GetPortinWithChangeMSISDN GetPortinWithChangeMSISDN `xml:"request"`
}

type GetPortinWithChangeMSISDNResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp   string `xml:"timestamp,attr"`
		Status      string `xml:"status,attr"`
		Type        string `xml:"type,attr"`
		Mdn         string `xml:"mdn"`
		Sim         string `xml:"sim"`
		Result      string `xml:"result"`
		Description string `xml:"description"`
	}
}
