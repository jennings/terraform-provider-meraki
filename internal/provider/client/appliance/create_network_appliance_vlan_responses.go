// Code generated by go-swagger; DO NOT EDIT.

package appliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateNetworkApplianceVlanReader is a Reader for the CreateNetworkApplianceVlan structure.
type CreateNetworkApplianceVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkApplianceVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkApplianceVlanCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkApplianceVlanCreated creates a CreateNetworkApplianceVlanCreated with default headers values
func NewCreateNetworkApplianceVlanCreated() *CreateNetworkApplianceVlanCreated {
	return &CreateNetworkApplianceVlanCreated{}
}

/* CreateNetworkApplianceVlanCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkApplianceVlanCreated struct {
	Payload interface{}
}

func (o *CreateNetworkApplianceVlanCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/appliance/vlans][%d] createNetworkApplianceVlanCreated  %+v", 201, o.Payload)
}
func (o *CreateNetworkApplianceVlanCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkApplianceVlanCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateNetworkApplianceVlanBody create network appliance vlan body
// Example: {"applianceIp":"192.168.1.2","groupPolicyId":"101","id":"1234","name":"My VLAN","subnet":"192.168.1.0/24"}
swagger:model CreateNetworkApplianceVlanBody
*/
type CreateNetworkApplianceVlanBody struct {

	// The local IP of the appliance on the VLAN
	ApplianceIP string `json:"applianceIp,omitempty"`

	// The id of the desired group policy to apply to the VLAN
	GroupPolicyID string `json:"groupPolicyId,omitempty"`

	// The VLAN ID of the new VLAN (must be between 1 and 4094)
	// Required: true
	ID *string `json:"id"`

	// The name of the new VLAN
	// Required: true
	Name *string `json:"name"`

	// The subnet of the VLAN
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this create network appliance vlan body
func (o *CreateNetworkApplianceVlanBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceVlanBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkApplianceVlan"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkApplianceVlanBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkApplianceVlan"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network appliance vlan body based on context it is used
func (o *CreateNetworkApplianceVlanBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceVlanBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceVlanBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
