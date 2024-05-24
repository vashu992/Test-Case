package model

type GetValidatePortOutEligibility struct {
	Type                 string      `xml:"type,attr"`
	MDN                  string      `xml:"MDN" validate:"required,len=10,numeric"`
	SIM                  string      `xml:"SIM"`
	IMEI                 string      `xml:"IMEI"`
	Name                 string      `xml:"name"`
	OspAccountNumber     string      `xml:"ospAccountNumber" validate:"required,len=7,numeric"`
	OspAccountPassword   string      `xml:"ospAccountPassword"`
	OspSubscriberAddress E911Address `xml:"ospSubscriberAddress"`
}

type GetValidatePortOutEligibilityRequest struct {
	Session                       AuthSession                   `xml:"session"`
	GetValidatePortOutEligibility GetValidatePortOutEligibility `xml:"request"`
}

type GetValidatePortOutEligibilityResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		AttrStatus string `xml:"status,attr"`
		Timestamp  string `xml:"timestamp,attr"`
		Type       string `xml:"type,attr"`
		Result     string `xml:"result"`
		ResultMsg  string `xml:"resultMsg"`
		Status     string `xml:"status"`
	}
}
