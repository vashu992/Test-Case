package pwg

import (
	"context"
	"encoding/xml"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
)

// GetCoverage2 receives XML request body, validates it, and responds with XML response.
func (c *Client) GetCoverage2(ctx context.Context, req *model.GetCoverage2Request) (*model.GetCoverage2Response, error) {

	req.Session = GenerateAuthSession(c.AuthConfig)

	res, err := c.SendRequest(ctx, req)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	coverage2Response := model.GetCoverage2Response{}

	err = xml.Unmarshal([]byte(res), &coverage2Response)
	if err != nil {
		c.Logger.Error(err.Error())
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
			Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	return &coverage2Response, nil
}
