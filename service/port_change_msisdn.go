package pwg

import (
	"context"
	"encoding/xml"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
)

// GetPortinWithChangeMSISDN receives XML request body, validates it, and responds with XML response.
func (c *Client) GetPortinWithChangeMSISDN(ctx context.Context, req *model.GetPortinWithChangeMSISDNRequest) (*model.GetPortinWithChangeMSISDNResponse, error) {

	req.Session = GenerateAuthSession(c.AuthConfig)

	res, err := c.SendRequest(ctx, req)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	portinWithChangeMSISDNResponse := model.GetPortinWithChangeMSISDNResponse{}

	err = xml.Unmarshal([]byte(res), &portinWithChangeMSISDNResponse)
	if err != nil {
		c.Logger.Error(err.Error())
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
			Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	return &portinWithChangeMSISDNResponse, nil
}
