package model

type GetChangePWGCostPlan struct {
	Type        string `xml:"type,attr"`
	Mdn         string `xml:"mdn" validate:"required,len=10,numeric"`
	Billingcode string `xml:"billingcode" validate:"required"`
}

type GetChangePWGCostPlanRequest struct {
	Session              AuthSession          `xml:"session"`
	GetChangePWGCostPlan GetChangePWGCostPlan `xml:"request"`
}
type GetChangePWGCostPlanResponse struct {
	Credentials struct {
		ReferenceNumber string `xml:"referenceNumber"`
		ReturnURL       string `xml:"returnURL"`
	}
	WholeSaleOrderResponse struct {
		StatusCode  string `xml:"statusCode"`
		Description string `xml:"description"`
	}
}
