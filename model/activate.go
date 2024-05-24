package model

type GetActivate struct {
	Esn         string      `xml:"esn" validate:"required,max=25,alphanum"`
	PlanId      string      `xml:"planId" validate:"required"`
	Language    string      `xml:"language"`
	Zip         string      `xml:"zip" validate:"required,len=5"`
	BillingCode string      `xml:"BillingCode" validate:"max=10,min=4"`
	E911Address E911Address `xml:"E911ADDRESS"`
}

type GetActivateRequest struct {
	Session     AuthSession `xml:"session"`
	GetActivate GetActivate `xml:"request"`
}

type GetActivateResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Status         string `xml:"status,attr"`
		Timestamp      string `xml:"timestamp,attr"`
		Type           string `xml:"type,attr"`
		Mdn            string `xml:"mdn"`
		Esn            string `xml:"esn"`
		CustomerID     string `xml:"CustomerID"`
		SubscriptionID string `xml:"SubscriptionID"`
		Warning        string `xml:"warning"`
	}
}
