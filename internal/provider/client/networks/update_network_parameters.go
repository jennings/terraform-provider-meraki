// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewUpdateNetworkParams creates a new UpdateNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkParams() *UpdateNetworkParams {
	return &UpdateNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkParamsWithTimeout creates a new UpdateNetworkParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkParamsWithTimeout(timeout time.Duration) *UpdateNetworkParams {
	return &UpdateNetworkParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkParamsWithContext creates a new UpdateNetworkParams object
// with the ability to set a context for a request.
func NewUpdateNetworkParamsWithContext(ctx context.Context) *UpdateNetworkParams {
	return &UpdateNetworkParams{
		Context: ctx,
	}
}

// NewUpdateNetworkParamsWithHTTPClient creates a new UpdateNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkParamsWithHTTPClient(client *http.Client) *UpdateNetworkParams {
	return &UpdateNetworkParams{
		HTTPClient: client,
	}
}

/* UpdateNetworkParams contains all the parameters to send to the API endpoint
   for the update network operation.

   Typically these are written to a http.Request.
*/
type UpdateNetworkParams struct {

	// NetworkID.
	NetworkID string

	// UpdateNetwork.
	UpdateNetwork UpdateNetworkBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkParams) WithDefaults() *UpdateNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network params
func (o *UpdateNetworkParams) WithTimeout(timeout time.Duration) *UpdateNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network params
func (o *UpdateNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network params
func (o *UpdateNetworkParams) WithContext(ctx context.Context) *UpdateNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network params
func (o *UpdateNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network params
func (o *UpdateNetworkParams) WithHTTPClient(client *http.Client) *UpdateNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network params
func (o *UpdateNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network params
func (o *UpdateNetworkParams) WithNetworkID(networkID string) *UpdateNetworkParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network params
func (o *UpdateNetworkParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetwork adds the updateNetwork to the update network params
func (o *UpdateNetworkParams) WithUpdateNetwork(updateNetwork UpdateNetworkBody) *UpdateNetworkParams {
	o.SetUpdateNetwork(updateNetwork)
	return o
}

// SetUpdateNetwork adds the updateNetwork to the update network params
func (o *UpdateNetworkParams) SetUpdateNetwork(updateNetwork UpdateNetworkBody) {
	o.UpdateNetwork = updateNetwork
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetwork); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
