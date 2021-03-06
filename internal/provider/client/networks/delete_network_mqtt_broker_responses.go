// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkMqttBrokerReader is a Reader for the DeleteNetworkMqttBroker structure.
type DeleteNetworkMqttBrokerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkMqttBrokerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkMqttBrokerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkMqttBrokerNoContent creates a DeleteNetworkMqttBrokerNoContent with default headers values
func NewDeleteNetworkMqttBrokerNoContent() *DeleteNetworkMqttBrokerNoContent {
	return &DeleteNetworkMqttBrokerNoContent{}
}

/* DeleteNetworkMqttBrokerNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkMqttBrokerNoContent struct {
}

func (o *DeleteNetworkMqttBrokerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/mqttBrokers/{mqttBrokerId}][%d] deleteNetworkMqttBrokerNoContent ", 204)
}

func (o *DeleteNetworkMqttBrokerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
