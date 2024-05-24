package model

type GetServiceInformationRequest struct {
	Session   AuthSession `xml:"session"`
	GetMdnEsn GetMdnEsn   `xml:"request"`
}

type GetServiceInformationResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		AttrStatus   string `xml:"status,attr"`
		Timestamp    string `xml:"timestamp,attr"`
		Type         string `xml:"type,attr"`
		Mdn          string `xml:"mdn"`
		Sim          string `xml:"sim"`
		IMEI         string `xml:"IMEI"`
		Status       string `xml:"status"`
		BillCycleDay string `xml:"billCycleDay"`
		Plan         struct {
			Plan          string `xml:"plan"`
			EffectiveDate string `xml:"effectiveDate"`
		}
		Socs []struct {
			Soc           string `xml:"soc"`
			EffectiveDate string `xml:"effectiveDate"`
		}
		BALANCEDETAIL struct {
			SUBSCRIBERSTATE   string `xml:"SUBSCRIBERSTATE"`
			HOTLINENUMBER     string `xml:"HOTLINENUMBER"`
			HOTLINECHARGEABLE string `xml:"HOTLINECHARGEABLE"`
			PURCHASE          []struct {
				PURCHASEID string `xml:"PURCHASEID"`
				TARIFFNAME string `xml:"TARIFFNAME"`
				PLANCODE   string `xml:"PLANCODE"`
				BALANCES   struct {
					SUBSCRIPTIONID string `xml:"SUBSCRIPTIONID"`
					UOM            string `xml:"UOM"`
					BALANCE        string `xml:"BALANCE"`
					VALIDFROM      string `xml:"VALIDFROM"`
					VALIDTO        string `xml:"VALIDTO"`
				}
			}
			TOTALBALANCE struct {
				TALK string `xml:"TALK"`
				TEXT string `xml:"TEXT"`
				DATA string `xml:"DATA"`
			}
		}
		StatusCode  string `xml:"statusCode"`
		Description string `xml:"description"`
	}
}
