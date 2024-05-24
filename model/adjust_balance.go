package model

type GetAdjustBalanceSingle struct {
	Mdn            string `xml:"mdn" validate:"required,len=10,numeric"`
	SubscriptionId string `xml:"subscriptionId" validate:"required,numeric"`
	Uom            string `xml:"uom"`
	Amount         string `xml:"amount"`
	ExpiryDate     string `xml:"expiryDate"`
}

type GetAdjustBalanceSingleRequest struct {
	Session                AuthSession            `xml:"session"`
	GetAdjustBalanceSingle GetAdjustBalanceSingle `xml:"request"`
}

type Subscription struct {
	Text   string `xml:",chardata"`
	Uom    string `xml:"uom"`
	Amount string `xml:"amount"`
}

type Subscriptions struct {
	SubscriptionId string `xml:"subscriptionId" validate:"required,numeric"`
	ExpiryDate     string `xml:"expiryDate"`
	Subscription   []Subscription
}

type GetAdjustMultipleBalancePayload struct {
	Mdn           string `xml:"mdn" validate:"required,len=10,numeric"`
	Subscriptions Subscriptions
}
type GetAdjustMultipleBalanceRequest struct {
	Session AuthSession                     `xml:"session"`
	Request GetAdjustMultipleBalancePayload `xml:"request"`
}
type GetAdjustBalanceSingleResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		MDN       string `xml:"MDN"`
	}
}

type GetAdjustMultipleBalanceResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		MDN       string `xml:"MDN"`
	}
}
