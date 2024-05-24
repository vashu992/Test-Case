package model

type GetRemoveLine struct {
	Type             string           `xml:"type,attr"`
	AddNewLineShared AddNewLineShared `xml:"RemoveLineShared"`
}

type GetRemoveLineRequest struct {
	Session       AuthSession   `xml:"session"`
	GetRemoveLine GetRemoveLine `xml:"request"`
}

type GetRemoveLineResponse struct {
	ResponseSession ResponseSession `xml:"session"`
	Response        struct {
		Timestamp string `xml:"timestamp,attr"`
		Status    string `xml:"status,attr"`
		Type      string `xml:"type,attr"`
		MDN       string `xml:"MDN"`
	}
}
