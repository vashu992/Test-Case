package model

type GetQueryHLRRequest struct {
	Session   AuthSession `xml:"session"`
	GetMdnSim GetMdnSim   `xml:"request"`
}

type GetQueryHLRResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Status    string `xml:"status,attr"`
		Timestamp string `xml:"timestamp,attr"`
		Type      string `xml:"type,attr"`
		Mdn       string `xml:"mdn"`
		Sim       string `xml:"sim"`
		IMEI      string `xml:"IMEI"`
		IMSI      string `xml:"IMSI"`
		SIMSTATUS string `xml:"SIMSTATUS"`
		Socs      []struct {
			Soc string `xml:"soc"`
		}
		Apn []struct {
			Name  string `xml:"name"`
			Value string `xml:"value"`
		}
		MSSTATUS    string `xml:"MS_STATUS"`
		StatusCode  string `xml:"statusCode"`
		Description string `xml:"description"`
	}
}
