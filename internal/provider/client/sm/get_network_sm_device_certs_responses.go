// Code generated by go-swagger; DO NOT EDIT.

package sm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSmDeviceCertsReader is a Reader for the GetNetworkSmDeviceCerts structure.
type GetNetworkSmDeviceCertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmDeviceCertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmDeviceCertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSmDeviceCertsOK creates a GetNetworkSmDeviceCertsOK with default headers values
func NewGetNetworkSmDeviceCertsOK() *GetNetworkSmDeviceCertsOK {
	return &GetNetworkSmDeviceCertsOK{}
}

/* GetNetworkSmDeviceCertsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmDeviceCertsOK struct {
	Payload []interface{}
}

func (o *GetNetworkSmDeviceCertsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/certs][%d] getNetworkSmDeviceCertsOK  %+v", 200, o.Payload)
}
func (o *GetNetworkSmDeviceCertsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkSmDeviceCertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
