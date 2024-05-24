package model

type GetChangeVoicemailLanguage struct {
	Type     string `xml:"type,attr"`
	Mdn      string `xml:"mdn" validate:"required,len=10,numeric"`
	Sim      string `xml:"sim"`
	Language string `xml:"language" validate:"required"`
}

type GetChangeVoicemailLanguageRequest struct {
	Session                    AuthSession                `xml:"session"`
	GetChangeVoicemailLanguage GetChangeVoicemailLanguage `xml:"request"`
}

type GetChangeVoicemailLanguageResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		Mdn       string `xml:"mdn"`
		Sim       string `xml:"sim"`
	}
}
