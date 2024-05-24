package model

type Address struct {
	StreetNumber    string `xml:"streetNumber"`
	StreetName      string `xml:"streetName"`
	StreetDirection string `xml:"streetDirection"`
	Line2           string `xml:"line2"`
	City            string `xml:"city"`
	State           string `xml:"state"`
	Zip             string `xml:"zip" validate:"required,max=5"`
}

type E911ADDRESS struct {
	STREET1 string `xml:"STREET1" validate:"required"`
	STREET2 string `xml:"STREET2"`
	CITY    string `xml:"CITY" validate:"required"`
	STATE   string `xml:"STATE" validate:"required"`
	ZIP     string `xml:"ZIP" validate:"required,len=5,numeric"`
}

type E911Address struct {
	Street1 string `xml:"street1" validate:"required"`
	Street2 string `xml:"street2"`
	City    string `xml:"city" validate:"required"`
	State   string `xml:"state" validate:"required"`
	Zip     string `xml:"zip" validate:"required,len=5,numeric"`
}

type Billing struct {
	FirstName string  `xml:"firstName" validate:"required,max=50"`
	LastName  string  `xml:"lastName" validate:"required,max=50"`
	Address   Address `xml:"address"`
}

type GetMdnSim struct {
	Type string `xml:"type,attr"`
	Mdn  string `xml:"mdn"`
	Sim  string `xml:"sim"`
}

type Hotline struct {
	Service string `xml:"service"`
}

type GetMdnEsn struct {
	Type string `xml:"type,attr"`
	Esn  string `xml:"esn"`
	Mdn  string `xml:"mdn"`
}

type Line struct { // 4
	MDN string `xml:"MDN" validate:"required,len=10,numeric"`
	SIM string `xml:"SIM"`
}

type LineDetails struct {
	Line Line `xml:"Line"`
}

type AddNewLineShared struct {
	ParentMDN   string      `xml:"parent-MDN" validate:"required,len=10,numeric"`
	ParentSIM   string      `xml:"parent-SIM"`
	LineDetails LineDetails `xml:"LineDetails"`
}

type ResponseSession struct {
	Clec struct {
		ID string `xml:"id"`
	}
	Timestamp string `xml:"timestamp"`
}

type WrappedError struct {
	Status  int64
	Message string
	Err     error
}
