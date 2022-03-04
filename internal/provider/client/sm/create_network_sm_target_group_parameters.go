// Code generated by go-swagger; DO NOT EDIT.

package sm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCreateNetworkSmTargetGroupParams creates a new CreateNetworkSmTargetGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkSmTargetGroupParams() *CreateNetworkSmTargetGroupParams {
	return &CreateNetworkSmTargetGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkSmTargetGroupParamsWithTimeout creates a new CreateNetworkSmTargetGroupParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkSmTargetGroupParamsWithTimeout(timeout time.Duration) *CreateNetworkSmTargetGroupParams {
	return &CreateNetworkSmTargetGroupParams{
		timeout: timeout,
	}
}

// NewCreateNetworkSmTargetGroupParamsWithContext creates a new CreateNetworkSmTargetGroupParams object
// with the ability to set a context for a request.
func NewCreateNetworkSmTargetGroupParamsWithContext(ctx context.Context) *CreateNetworkSmTargetGroupParams {
	return &CreateNetworkSmTargetGroupParams{
		Context: ctx,
	}
}

// NewCreateNetworkSmTargetGroupParamsWithHTTPClient creates a new CreateNetworkSmTargetGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkSmTargetGroupParamsWithHTTPClient(client *http.Client) *CreateNetworkSmTargetGroupParams {
	return &CreateNetworkSmTargetGroupParams{
		HTTPClient: client,
	}
}

/* CreateNetworkSmTargetGroupParams contains all the parameters to send to the API endpoint
   for the create network sm target group operation.

   Typically these are written to a http.Request.
*/
type CreateNetworkSmTargetGroupParams struct {

	// CreateNetworkSmTargetGroup.
	CreateNetworkSmTargetGroup CreateNetworkSmTargetGroupBody

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network sm target group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSmTargetGroupParams) WithDefaults() *CreateNetworkSmTargetGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network sm target group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSmTargetGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network sm target group params
func (o *CreateNetworkSmTargetGroupParams) WithTimeout(timeout time.Duration) *CreateNetworkSmTargetGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network sm target group params
func (o *CreateNetworkSmTargetGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network sm target group params
func (o *CreateNetworkSmTargetGroupParams) WithContext(ctx context.Context) *CreateNetworkSmTargetGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network sm target group params
func (o *CreateNetworkSmTargetGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network sm target group params
func (o *CreateNetworkSmTargetGroupParams) WithHTTPClient(client *http.Client) *CreateNetworkSmTargetGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network sm target group params
func (o *CreateNetworkSmTargetGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkSmTargetGroup adds the createNetworkSmTargetGroup to the create network sm target group params
func (o *CreateNetworkSmTargetGroupParams) WithCreateNetworkSmTargetGroup(createNetworkSmTargetGroup CreateNetworkSmTargetGroupBody) *CreateNetworkSmTargetGroupParams {
	o.SetCreateNetworkSmTargetGroup(createNetworkSmTargetGroup)
	return o
}

// SetCreateNetworkSmTargetGroup adds the createNetworkSmTargetGroup to the create network sm target group params
func (o *CreateNetworkSmTargetGroupParams) SetCreateNetworkSmTargetGroup(createNetworkSmTargetGroup CreateNetworkSmTargetGroupBody) {
	o.CreateNetworkSmTargetGroup = createNetworkSmTargetGroup
}

// WithNetworkID adds the networkID to the create network sm target group params
func (o *CreateNetworkSmTargetGroupParams) WithNetworkID(networkID string) *CreateNetworkSmTargetGroupParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network sm target group params
func (o *CreateNetworkSmTargetGroupParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkSmTargetGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateNetworkSmTargetGroup); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
