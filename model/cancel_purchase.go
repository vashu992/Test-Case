package model

type GetCancelPurchase struct {
	Type       string `xml:"type,attr"`
	Mdn        string `xml:"mdn" validate:"required,len=10"`
	PurchaseId string `xml:"purchaseId" validate:"required,max=35"`
}

type GetCancelPurchaseRequest struct {
	Session           AuthSession       `xml:"session"`
	GetCancelPurchase GetCancelPurchase `xml:"request"`
}

type GetCancelPurchaseResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp  string `xml:"timestamp,attr"`
		Status     string `xml:"status,attr"`
		Type       string `xml:"type,attr"`
		PurchaseId string `xml:"purchaseId"`
	}
}
