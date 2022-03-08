// Code generated by go-swagger; DO NOT EDIT.

package appliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkApplianceContentFilteringReader is a Reader for the GetNetworkApplianceContentFiltering structure.
type GetNetworkApplianceContentFilteringReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkApplianceContentFilteringReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkApplianceContentFilteringOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkApplianceContentFilteringOK creates a GetNetworkApplianceContentFilteringOK with default headers values
func NewGetNetworkApplianceContentFilteringOK() *GetNetworkApplianceContentFilteringOK {
	return &GetNetworkApplianceContentFilteringOK{}
}

/* GetNetworkApplianceContentFilteringOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkApplianceContentFilteringOK struct {
	Payload interface{}
}

func (o *GetNetworkApplianceContentFilteringOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/contentFiltering][%d] getNetworkApplianceContentFilteringOK  %+v", 200, o.Payload)
}
func (o *GetNetworkApplianceContentFilteringOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkApplianceContentFilteringOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
