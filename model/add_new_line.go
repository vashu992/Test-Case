package model

type GetAddNewLine struct {
	Type             string           `xml:"type,attr"`
	AddNewLineShared AddNewLineShared `xml:"AddNewLineShared"`
}

type GetAddNewLineRequest struct {
	Session       AuthSession   `xml:"session"`
	GetAddNewLine GetAddNewLine `xml:"request"`
}

type GetAddNewLineResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		MDN       string `xml:"MDN"`
	}
}
