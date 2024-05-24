package model

type GetPurchasePlan struct {
	Type        string      `xml:"type,attr"`
	Mdn         string      `xml:"mdn"`
	PlanId      string      `xml:"planId"`
	E911ADDRESS E911ADDRESS `xml:"E911ADDRESS"`
}
type GetPurchasePlanRequest struct {
	Session         AuthSession     `xml:"session"`
	GetPurchasePlan GetPurchasePlan `xml:"request"`
}

type GetPurchasePlanResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp  string `xml:"timestamp,attr"`
		Status     string `xml:"status,attr"`
		Type       string `xml:"type,attr"`
		PurchaseId string `xml:"purchaseId"`
		Warning    string `xml:"warning"`
	}
}
