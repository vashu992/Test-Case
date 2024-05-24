package model

type GetQueryUsageRequest struct {
	Session   AuthSession `xml:"session"`
	GetMdnEsn GetMdnEsn   `xml:"request"`
}

type GetQueryUsageResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Status        string `xml:"status,attr"`
		Timestamp     string `xml:"timestamp,attr"`
		Type          string `xml:"type,attr"`
		MDN           string `xml:"MDN" validate:"required,len=10,numeric"`
		SIM           string `xml:"SIM"`
		ACCOUNTSTATUS string `xml:"ACCOUNTSTATUS"`
		SOC           string `xml:"SOC"`
		TYPE          string `xml:"TYPE"`
		LIMIT         string `xml:"LIMIT"`
		USED          string `xml:"USED"`
		USAGESTATUS   string `xml:"USAGESTATUS"`
		STATUSCODE    string `xml:"STATUSCODE"`
		DESCRIPTION   string `xml:"DESCRIPTION"`
	}
}
