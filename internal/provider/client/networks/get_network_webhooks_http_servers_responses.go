// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkWebhooksHTTPServersReader is a Reader for the GetNetworkWebhooksHTTPServers structure.
type GetNetworkWebhooksHTTPServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWebhooksHTTPServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWebhooksHTTPServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWebhooksHTTPServersOK creates a GetNetworkWebhooksHTTPServersOK with default headers values
func NewGetNetworkWebhooksHTTPServersOK() *GetNetworkWebhooksHTTPServersOK {
	return &GetNetworkWebhooksHTTPServersOK{}
}

/* GetNetworkWebhooksHTTPServersOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWebhooksHTTPServersOK struct {
	Payload []interface{}
}

func (o *GetNetworkWebhooksHTTPServersOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/webhooks/httpServers][%d] getNetworkWebhooksHttpServersOK  %+v", 200, o.Payload)
}
func (o *GetNetworkWebhooksHTTPServersOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkWebhooksHTTPServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
