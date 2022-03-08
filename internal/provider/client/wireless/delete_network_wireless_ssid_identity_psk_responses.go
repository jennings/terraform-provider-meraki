// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkWirelessSsidIdentityPskReader is a Reader for the DeleteNetworkWirelessSsidIdentityPsk structure.
type DeleteNetworkWirelessSsidIdentityPskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkWirelessSsidIdentityPskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkWirelessSsidIdentityPskNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkWirelessSsidIdentityPskNoContent creates a DeleteNetworkWirelessSsidIdentityPskNoContent with default headers values
func NewDeleteNetworkWirelessSsidIdentityPskNoContent() *DeleteNetworkWirelessSsidIdentityPskNoContent {
	return &DeleteNetworkWirelessSsidIdentityPskNoContent{}
}

/* DeleteNetworkWirelessSsidIdentityPskNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkWirelessSsidIdentityPskNoContent struct {
}

func (o *DeleteNetworkWirelessSsidIdentityPskNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}][%d] deleteNetworkWirelessSsidIdentityPskNoContent ", 204)
}

func (o *DeleteNetworkWirelessSsidIdentityPskNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
