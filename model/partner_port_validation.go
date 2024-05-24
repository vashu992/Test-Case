package model

type GetPartnerPortOutValidation struct {
	Type        string `xml:"type,attr"`
	Mdn         string `xml:"mdn" validate:"required,len=10,numeric"`
	Sim         string `xml:"sim" validate:"required,max=25"`
	MessageCode string `xml:"MessageCode" validate:"required,alphanum"`
	Description string `xml:"Description" validate:"required,alphanum"`
}
type GetPartnerPortOutValidationRequest struct {
	Session                     AuthSession                 `xml:"session"`
	GetPartnerPortOutValidation GetPartnerPortOutValidation `xml:"request"`
}

type GetPartnerPortOutValidationResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Status      string `xml:"status,attr"`
		Timestamp   string `xml:"timestamp,attr"`
		Type        string `xml:"type,attr"`
		Mdn         string `xml:"mdn"`
		Sim         string `xml:"sim"`
		StatusCode  string `xml:"statusCode"`
		Description string `xml:"description"`
	}
}
