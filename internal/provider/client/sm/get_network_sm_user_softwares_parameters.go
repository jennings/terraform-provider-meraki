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

// NewGetNetworkSmUserSoftwaresParams creates a new GetNetworkSmUserSoftwaresParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSmUserSoftwaresParams() *GetNetworkSmUserSoftwaresParams {
	return &GetNetworkSmUserSoftwaresParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSmUserSoftwaresParamsWithTimeout creates a new GetNetworkSmUserSoftwaresParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSmUserSoftwaresParamsWithTimeout(timeout time.Duration) *GetNetworkSmUserSoftwaresParams {
	return &GetNetworkSmUserSoftwaresParams{
		timeout: timeout,
	}
}

// NewGetNetworkSmUserSoftwaresParamsWithContext creates a new GetNetworkSmUserSoftwaresParams object
// with the ability to set a context for a request.
func NewGetNetworkSmUserSoftwaresParamsWithContext(ctx context.Context) *GetNetworkSmUserSoftwaresParams {
	return &GetNetworkSmUserSoftwaresParams{
		Context: ctx,
	}
}

// NewGetNetworkSmUserSoftwaresParamsWithHTTPClient creates a new GetNetworkSmUserSoftwaresParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSmUserSoftwaresParamsWithHTTPClient(client *http.Client) *GetNetworkSmUserSoftwaresParams {
	return &GetNetworkSmUserSoftwaresParams{
		HTTPClient: client,
	}
}

/* GetNetworkSmUserSoftwaresParams contains all the parameters to send to the API endpoint
   for the get network sm user softwares operation.

   Typically these are written to a http.Request.
*/
type GetNetworkSmUserSoftwaresParams struct {

	// NetworkID.
	NetworkID string

	// UserID.
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network sm user softwares params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmUserSoftwaresParams) WithDefaults() *GetNetworkSmUserSoftwaresParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network sm user softwares params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmUserSoftwaresParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network sm user softwares params
func (o *GetNetworkSmUserSoftwaresParams) WithTimeout(timeout time.Duration) *GetNetworkSmUserSoftwaresParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sm user softwares params
func (o *GetNetworkSmUserSoftwaresParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sm user softwares params
func (o *GetNetworkSmUserSoftwaresParams) WithContext(ctx context.Context) *GetNetworkSmUserSoftwaresParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sm user softwares params
func (o *GetNetworkSmUserSoftwaresParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sm user softwares params
func (o *GetNetworkSmUserSoftwaresParams) WithHTTPClient(client *http.Client) *GetNetworkSmUserSoftwaresParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sm user softwares params
func (o *GetNetworkSmUserSoftwaresParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network sm user softwares params
func (o *GetNetworkSmUserSoftwaresParams) WithNetworkID(networkID string) *GetNetworkSmUserSoftwaresParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sm user softwares params
func (o *GetNetworkSmUserSoftwaresParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUserID adds the userID to the get network sm user softwares params
func (o *GetNetworkSmUserSoftwaresParams) WithUserID(userID string) *GetNetworkSmUserSoftwaresParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get network sm user softwares params
func (o *GetNetworkSmUserSoftwaresParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSmUserSoftwaresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
