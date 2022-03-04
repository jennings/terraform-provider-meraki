// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateNetworkSwitchQosRuleReader is a Reader for the CreateNetworkSwitchQosRule structure.
type CreateNetworkSwitchQosRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSwitchQosRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSwitchQosRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkSwitchQosRuleCreated creates a CreateNetworkSwitchQosRuleCreated with default headers values
func NewCreateNetworkSwitchQosRuleCreated() *CreateNetworkSwitchQosRuleCreated {
	return &CreateNetworkSwitchQosRuleCreated{}
}

/* CreateNetworkSwitchQosRuleCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkSwitchQosRuleCreated struct {
	Payload interface{}
}

func (o *CreateNetworkSwitchQosRuleCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/qosRules][%d] createNetworkSwitchQosRuleCreated  %+v", 201, o.Payload)
}
func (o *CreateNetworkSwitchQosRuleCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSwitchQosRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateNetworkSwitchQosRuleBody create network switch qos rule body
// Example: {"dscp":0,"dstPort":3000,"dstPortRange":"3000-3100","protocol":"TCP","srcPort":2000,"srcPortRange":"70-80","vlan":100}
swagger:model CreateNetworkSwitchQosRuleBody
*/
type CreateNetworkSwitchQosRuleBody struct {

	// DSCP tag. Set this to -1 to trust incoming DSCP. Default value is 0
	Dscp int64 `json:"dscp,omitempty"`

	// The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPort int64 `json:"dstPort,omitempty"`

	// The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	DstPortRange string `json:"dstPortRange,omitempty"`

	// The protocol of the incoming packet. Can be one of "ANY", "TCP" or "UDP". Default value is "ANY"
	// Enum: [ANY TCP UDP]
	Protocol string `json:"protocol,omitempty"`

	// The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPort int64 `json:"srcPort,omitempty"`

	// The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	SrcPortRange string `json:"srcPortRange,omitempty"`

	// The VLAN of the incoming packet. A null value will match any VLAN.
	// Required: true
	Vlan *int64 `json:"vlan"`
}

// Validate validates this create network switch qos rule body
func (o *CreateNetworkSwitchQosRuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createNetworkSwitchQosRuleBodyTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ANY","TCP","UDP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkSwitchQosRuleBodyTypeProtocolPropEnum = append(createNetworkSwitchQosRuleBodyTypeProtocolPropEnum, v)
	}
}

const (

	// CreateNetworkSwitchQosRuleBodyProtocolANY captures enum value "ANY"
	CreateNetworkSwitchQosRuleBodyProtocolANY string = "ANY"

	// CreateNetworkSwitchQosRuleBodyProtocolTCP captures enum value "TCP"
	CreateNetworkSwitchQosRuleBodyProtocolTCP string = "TCP"

	// CreateNetworkSwitchQosRuleBodyProtocolUDP captures enum value "UDP"
	CreateNetworkSwitchQosRuleBodyProtocolUDP string = "UDP"
)

// prop value enum
func (o *CreateNetworkSwitchQosRuleBody) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkSwitchQosRuleBodyTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkSwitchQosRuleBody) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(o.Protocol) { // not required
		return nil
	}

	// value enum
	if err := o.validateProtocolEnum("createNetworkSwitchQosRule"+"."+"protocol", "body", o.Protocol); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchQosRuleBody) validateVlan(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchQosRule"+"."+"vlan", "body", o.Vlan); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network switch qos rule body based on context it is used
func (o *CreateNetworkSwitchQosRuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchQosRuleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchQosRuleBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchQosRuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
