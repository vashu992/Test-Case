package model

type GetDeviceSwapNotification struct {
	Mdn     string `xml:"mdn" validate:"required,len=10,numeric"`
	Newimei string `xml:"newimei" validate:"required,max=25"`
}

type GetDeviceSwapNotificationRequest struct {
	Session                   AuthSession               `xml:"session"`
	GetDeviceSwapNotification GetDeviceSwapNotification `xml:"request"`
}

type GetDeviceSwapNotificationResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		Mdn       string `xml:"mdn"`
		Imei      string `xml:"imei"`
	}
}
