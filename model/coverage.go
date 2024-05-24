package model

type CoverageRequest struct {
	Session AuthSession `xml:"session"`
	Request Coverage    `xml:"request"`
}

type Coverage struct {
	Carrier string `xml:"carrier" validate:"required"`
	Zip     string `xml:"zip" validate:"required,min=6,max=10"`
	Type    string `xml:"_type"`
}

// GetCoverageRequest contains data for GetCoverage method.
type GetCoverageRequest struct {
	CoverageRequest
}

type GetCoverageResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Status              string   `xml:"status,attr"`
		Timestamp           string   `xml:"timestamp,attr"`
		Type                string   `xml:"type,attr"`
		Zip                 string   `xml:"zip"`
		StatusCode          string   `xml:"statusCode"`
		CoverageQualityIden []string `xml:"coverageQualityIden"`
		Csa                 []string `xml:"csa"`
	}
}
