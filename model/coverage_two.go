package model

type GetCoverage2 struct {
	Type string `xml:"type,attr"`
	Zip  string `xml:"zip" validate:"required,max=5,numeric"`
}

type GetCoverage2Request struct {
	Session      AuthSession  `xml:"session"`
	GetCoverage2 GetCoverage2 `xml:"request"`
}

type GetCoverage2Response struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Status             string `xml:"status,attr"`
		Timestamp          string `xml:"timestamp,attr"`
		Type               string `xml:"type,attr"`
		Zip                string `xml:"zip"`
		StatusCode         string `xml:"statusCode"`
		ActivationEligible string `xml:"activationEligible"`
	}
}
