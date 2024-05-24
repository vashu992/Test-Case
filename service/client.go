package pwg

import (
	"bytes"
	"context"
	"encoding/xml"
	"io"
	"net/http"
	constants "propwg/common"
	util "propwg/common"

	"github.com/apex/log"
	"github.com/go-playground/validator/v10"
)

// Client represents the propwg client.
type Client struct {
	AuthConfig AuthConfig
	Logger     *log.Logger
	HTTPClient *http.Client
}

type AuthConfig struct {
	Url      string
	ClecId   string
	UserName string
	Token    string
	Pin      string
}

// NewClient creates a new propwg client.
func NewClient(authconfig AuthConfig, logger *log.Logger, httpClient *http.Client) *Client {
	return &Client{
		AuthConfig: authconfig,
		Logger:     logger,
		HTTPClient: httpClient,
	}
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func (c *Client) SendRequest(ctx context.Context, postData interface{}) ([]byte, error) {
	err := validate.Struct(postData)
	if err != nil {
		pwgErr := &util.WrappedError{
			Code:    constants.REQUEST_VALIDATION_EXCPETION_CODE,
			Message: constants.REQUEST_VALIDATION_EXCPETION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	marshalledData, err := xml.Marshal(postData)
	if err != nil {
		pwgErr := &util.WrappedError{
			Code:    constants.REQUEST_MARSHAL_EXCEPTION_CODE,
			Message: constants.REQUEST_MARSHAL_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.AuthConfig.Url, bytes.NewReader(marshalledData))
	if err != nil {
		pwgErr := &util.WrappedError{
			Code:    constants.REQUEST_GENERATION_EXCEPTION_CODE,
			Message: constants.REQUEST_GENERATION_EXCEPTION_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	req.Header.Set("Content-Type", "text/xml")

	// Post the request to endpoint
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		pwgErr := &util.WrappedError{
			Code:    constants.REQUEST_FAILED_CODE,
			Message: constants.REQUEST_FAILED_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_ERROR_CODE,
			Message: constants.RESPONSE_ERROR_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		pwgErr := &util.WrappedError{
			Code:    constants.RESPONSE_IO_ERROR_CODE,
			Message: constants.RESPONSE_IO_ERROR_INFO,
			Err:     err,
		}
		return nil, pwgErr
	}
	return bodyBytes, nil
}
