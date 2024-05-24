package pwg

import (
	"context"
	"encoding/xml"
	constants "propwg/common"
	util "propwg/common"
	"propwg/model"
)

// GetBalanceInformation receives XML request body, validates it, and responds with XML response.
func (c *Client) GetBalanceInformation(ctx context.Context, req *model.GetBalanceInformationRequest) (*model.GetBalanceInformationResponse, error) {

	req.Session = GenerateAuthSession(c.AuthConfig)

	res, err := c.SendRequest(ctx, req)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	balanceInformationResponse := model.GetBalanceInformationResponse{}

	err = xml.Unmarshal([]byte(res), &balanceInformationResponse)
	if err != nil {
		c.Logger.Error(err.Error())
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
			Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	return &balanceInformationResponse, nil
}

// GetBalanceInformationWithoutCredit receives XML request body, validates it, and responds with XML response.
func (c *Client) GetBalanceInformationWithoutCredit(ctx context.Context, req *model.GetBalanceInformationRequest) (*model.GetBalanceInformationWithoutCreditResponse, error) {

	req.Session = GenerateAuthSession(c.AuthConfig)

	res, err := c.SendRequest(ctx, req)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	GetBalanceInformationWithoutCreditResponse := model.GetBalanceInformationWithoutCreditResponse{}

	err = xml.Unmarshal([]byte(res), &GetBalanceInformationWithoutCreditResponse)
	if err != nil {
		c.Logger.Error(err.Error())
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
			Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	return &GetBalanceInformationWithoutCreditResponse, nil
}

// GetBalanceInformationPendingBalance receives XML request body, validates it, and responds with XML response.
func (c *Client) GetBalanceInformationPendingBalance(ctx context.Context, req *model.GetBalanceInformationRequest) (*model.GetBalanceInformationPendingBalanceResponse, error) {

	req.Session = GenerateAuthSession(c.AuthConfig)

	res, err := c.SendRequest(ctx, req)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	GetBalanceInformationPendingBalanceResponse := model.GetBalanceInformationPendingBalanceResponse{}

	err = xml.Unmarshal([]byte(res), &GetBalanceInformationPendingBalanceResponse)
	if err != nil {
		c.Logger.Error(err.Error())
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_UNMARSHAL_EXCEPTION_CODE,
			Message: constants.RESPONSE_UNMARSHAL_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	return &GetBalanceInformationPendingBalanceResponse, nil
}
