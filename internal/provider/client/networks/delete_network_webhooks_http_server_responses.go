// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkWebhooksHTTPServerReader is a Reader for the DeleteNetworkWebhooksHTTPServer structure.
type DeleteNetworkWebhooksHTTPServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkWebhooksHTTPServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkWebhooksHTTPServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkWebhooksHTTPServerNoContent creates a DeleteNetworkWebhooksHTTPServerNoContent with default headers values
func NewDeleteNetworkWebhooksHTTPServerNoContent() *DeleteNetworkWebhooksHTTPServerNoContent {
	return &DeleteNetworkWebhooksHTTPServerNoContent{}
}

/* DeleteNetworkWebhooksHTTPServerNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkWebhooksHTTPServerNoContent struct {
}

func (o *DeleteNetworkWebhooksHTTPServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/webhooks/httpServers/{httpServerId}][%d] deleteNetworkWebhooksHttpServerNoContent ", 204)
}

func (o *DeleteNetworkWebhooksHTTPServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
