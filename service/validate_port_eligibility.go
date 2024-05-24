package pwg

import (
	"context"
	"encoding/xml"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
)

// GetValidatePortOutEligibility receives XML request body, validates it, and responds with XML response.
func (c *Client) GetValidatePortOutEligibility(ctx context.Context, req *model.GetValidatePortOutEligibilityRequest) (*model.GetValidatePortOutEligibilityResponse, error) {

	req.Session = GenerateAuthSession(c.AuthConfig)

	res, err := c.SendRequest(ctx, req)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	validatePortOutEligibilityResponse := model.GetValidatePortOutEligibilityResponse{}

	err = xml.Unmarshal([]byte(res), &validatePortOutEligibilityResponse)
	if err != nil {
		c.Logger.Error(err.Error())
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
			Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	return &validatePortOutEligibilityResponse, nil
}
