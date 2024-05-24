package pwg

import (
	"context"
	"encoding/xml"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
)

// GetActivate receives XML request body, validates it, and responds with XML response.
func (c *Client) GetActivate(ctx context.Context, req *model.GetActivateRequest) (*model.GetActivateResponse, error) {

	req.Session = GenerateAuthSession(c.AuthConfig)

	res, err := c.SendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	activateResponse := model.GetActivateResponse{}

	err = xml.Unmarshal([]byte(res), &activateResponse)
	if err != nil {
		c.Logger.Error(err.Error())
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
			Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	return &activateResponse, nil
}
