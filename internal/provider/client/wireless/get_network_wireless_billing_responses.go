// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkWirelessBillingReader is a Reader for the GetNetworkWirelessBilling structure.
type GetNetworkWirelessBillingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessBillingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessBillingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessBillingOK creates a GetNetworkWirelessBillingOK with default headers values
func NewGetNetworkWirelessBillingOK() *GetNetworkWirelessBillingOK {
	return &GetNetworkWirelessBillingOK{}
}

/* GetNetworkWirelessBillingOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessBillingOK struct {
	Payload interface{}
}

func (o *GetNetworkWirelessBillingOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/billing][%d] getNetworkWirelessBillingOK  %+v", 200, o.Payload)
}
func (o *GetNetworkWirelessBillingOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkWirelessBillingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
