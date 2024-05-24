package model

type GetBalanceInformation struct {
	Type           string `xml:"type,attr"`
	Mdn            string `xml:"mdn" validate:"required,len=10"`
	Esn            string `xml:"esn" validate:"max=25,alphanum"`
	PendingBalance string `xml:"pendingBalance"`
}

type GetBalanceInformationRequest struct {
	Session               AuthSession           `xml:"session"`
	GetBalanceInformation GetBalanceInformation `xml:"request"`
}

type GetBalanceInformationResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		AttrStatus    string `xml:"status,attr"`
		Timestamp     string `xml:"timestamp,attr"`
		Type          string `xml:"type,attr"`
		Mdn           string `xml:"mdn"`
		Sim           string `xml:"sim"`
		IMEI          string `xml:"IMEI"`
		Status        string `xml:"status"`
		Customerid    string `xml:"customerid"`
		OUTOFCREDIT   string `xml:"OUTOFCREDIT"`
		BillCycleDay  string `xml:"billCycleDay"`
		BALANCEDETAIL struct {
			HOTLINENUMBER     string `xml:"HOTLINENUMBER"`
			HOTLINECHARGEABLE string `xml:"HOTLINECHARGEABLE"`
			PURCHASE          struct {
				PURCHASEID string `xml:"PURCHASEID"`
				TARIFFNAME string `xml:"TARIFFNAME"`
				PLANCODE   string `xml:"PLANCODE"`
				BALANCES   []struct {
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

type GetBalanceInformationWithoutCreditResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		AttrStatus    string `xml:"status,attr"`
		Timestamp     string `xml:"timestamp,attr"`
		Type          string `xml:"type,attr"`
		Mdn           string `xml:"mdn"`
		Sim           string `xml:"sim"`
		IMEI          string `xml:"IMEI"`
		Status        string `xml:"status"`
		Customerid    string `xml:"customerid"`
		BillCycleDay  string `xml:"billCycleDay"`
		BALANCEDETAIL struct {
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

// GetBalanceInformationPendingBalanceResponse
type GetBalanceInformationPendingBalanceResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		AttrStatus    string `xml:"status,attr"`
		Timestamp     string `xml:"timestamp,attr"`
		Type          string `xml:"type,attr"`
		Mdn           string `xml:"mdn"`
		Sim           string `xml:"sim"`
		IMEI          string `xml:"IMEI"`
		Status        string `xml:"status"`
		MVNOID        string `xml:"MVNO_ID"`
		Customerid    string `xml:"customerid"`
		BillCycleDay  string `xml:"billCycleDay"`
		BALANCEDETAIL struct {
			HOTLINENUMBER     string `xml:"HOTLINENUMBER"`
			HOTLINECHARGEABLE string `xml:"HOTLINECHARGEABLE"`
			PURCHASE          []struct {
				PURCHASEID string `xml:"PURCHASEID"`
				TARIFFNAME string `xml:"TARIFFNAME"`
				PLANCODE   string `xml:"PLANCODE"`
				BALANCES   []struct {
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
