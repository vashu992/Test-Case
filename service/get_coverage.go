package pwg

import (
	"context"
	"encoding/xml"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
)

// GetCoverage receives XML request body, validates it, and responds with XML response.
func (c *Client) GetCoverage(ctx context.Context, req *model.GetCoverageRequest) (*model.GetCoverageResponse, error) {

	req.Session = GenerateAuthSession(c.AuthConfig)

	res, err := c.SendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	coverageResponse := model.GetCoverageResponse{}

	err = xml.Unmarshal([]byte(res), &coverageResponse)
	if err != nil {
		c.Logger.Error(err.Error())
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
			Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	return &coverageResponse, nil
}
