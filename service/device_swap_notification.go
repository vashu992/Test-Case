package pwg

import (
	"context"
	"encoding/xml"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
)

// GetDeviceSwapNotification receives XML request body, validates it, and responds with XML response.
func (c *Client) GetDeviceSwapNotification(ctx context.Context, req *model.GetDeviceSwapNotificationRequest) (*model.GetDeviceSwapNotificationResponse, error) {

	req.Session = GenerateAuthSession(c.AuthConfig)

	res, err := c.SendRequest(ctx, req)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	deviceSwapNotificationResponse := model.GetDeviceSwapNotificationResponse{}

	err = xml.Unmarshal([]byte(res), &deviceSwapNotificationResponse)
	if err != nil {
		c.Logger.Error(err.Error())
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
			Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	return &deviceSwapNotificationResponse, nil
}
