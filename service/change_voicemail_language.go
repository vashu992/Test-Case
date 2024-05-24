package pwg

import (
	"context"
	"encoding/xml"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
)

// GetChangeVoicemailLanguage receives XML request body, validates it, and responds with XML response.
func (c *Client) GetChangeVoicemailLanguage(ctx context.Context, req *model.GetChangeVoicemailLanguageRequest) (*model.GetChangeVoicemailLanguageResponse, error) {

	req.Session = GenerateAuthSession(c.AuthConfig)

	res, err := c.SendRequest(ctx, req)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	changeVoicemailLanguageResponse := model.GetChangeVoicemailLanguageResponse{}

	err = xml.Unmarshal([]byte(res), &changeVoicemailLanguageResponse)
	if err != nil {
		c.Logger.Error(err.Error())
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
			Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	return &changeVoicemailLanguageResponse, nil
}
