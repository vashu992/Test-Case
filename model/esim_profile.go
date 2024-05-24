package model

type GetESIMProfile struct {
	Sim       string `xml:"sim"`
	ESimBrand string `xml:"eSimBrand" validate:"required"`
}

type GetESIMProfileRequest struct {
	Session        AuthSession    `xml:"session"`
	GetESIMProfile GetESIMProfile `xml:"request"`
}

type GetESIMProfileResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Status         string `xml:"status,attr"`
		Timestamp      string `xml:"timestamp,attr"`
		Type           string `xml:"type,attr"`
		StatusCode     string `xml:"statusCode"`
		Description    string `xml:"description"`
		Iccid          string `xml:"iccid"`
		ActivationCode string `xml:"activationCode"`
		ProfileType    string `xml:"profileType"`
		Lastmodified   string `xml:"lastmodified"`
		MatchingId     string `xml:"matchingId"`
		ProfileState   string `xml:"profileState"`
		ESimBrand      string `xml:"eSimBrand"`
	}
}
