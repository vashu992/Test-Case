package model

type GetQuerySIM struct {
	Esn string `xml:"esn"`
}

type GetQuerySIMRequest struct {
	Session     AuthSession `xml:"session"`
	GetQuerySIM GetQuerySIM `xml:"request"`
}

type GetQuerySIMResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Status             string `xml:"status,attr"`
		Timestamp          string `xml:"timestamp,attr"`
		Type               string `xml:"type,attr"`
		SIM                string `xml:"SIM"`
		ACTIVATIONELIGIBLE string `xml:"ACTIVATIONELIGIBLE"`
		WFCEligible        string `xml:"WFCEligible"`
		PUK1               string `xml:"PUK1"`
		PUK2               string `xml:"PUK2"`
		CREATEDATE         string `xml:"CREATEDATE"`
		EXPIRATIONDATE     string `xml:"EXPIRATIONDATE"`
		ICCIDSTATUS        string `xml:"ICCIDSTATUS"`
		StatusCode         string `xml:"statusCode"`
		Description        string `xml:"description"`
	}
}
