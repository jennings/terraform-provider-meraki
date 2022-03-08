// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewUpdateOrganizationActionBatchParams creates a new UpdateOrganizationActionBatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationActionBatchParams() *UpdateOrganizationActionBatchParams {
	return &UpdateOrganizationActionBatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationActionBatchParamsWithTimeout creates a new UpdateOrganizationActionBatchParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationActionBatchParamsWithTimeout(timeout time.Duration) *UpdateOrganizationActionBatchParams {
	return &UpdateOrganizationActionBatchParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationActionBatchParamsWithContext creates a new UpdateOrganizationActionBatchParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationActionBatchParamsWithContext(ctx context.Context) *UpdateOrganizationActionBatchParams {
	return &UpdateOrganizationActionBatchParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationActionBatchParamsWithHTTPClient creates a new UpdateOrganizationActionBatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationActionBatchParamsWithHTTPClient(client *http.Client) *UpdateOrganizationActionBatchParams {
	return &UpdateOrganizationActionBatchParams{
		HTTPClient: client,
	}
}

/* UpdateOrganizationActionBatchParams contains all the parameters to send to the API endpoint
   for the update organization action batch operation.

   Typically these are written to a http.Request.
*/
type UpdateOrganizationActionBatchParams struct {

	// ActionBatchID.
	ActionBatchID string

	// OrganizationID.
	OrganizationID string

	// UpdateOrganizationActionBatch.
	UpdateOrganizationActionBatch UpdateOrganizationActionBatchBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization action batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationActionBatchParams) WithDefaults() *UpdateOrganizationActionBatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization action batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationActionBatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) WithTimeout(timeout time.Duration) *UpdateOrganizationActionBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) WithContext(ctx context.Context) *UpdateOrganizationActionBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) WithHTTPClient(client *http.Client) *UpdateOrganizationActionBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionBatchID adds the actionBatchID to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) WithActionBatchID(actionBatchID string) *UpdateOrganizationActionBatchParams {
	o.SetActionBatchID(actionBatchID)
	return o
}

// SetActionBatchID adds the actionBatchId to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) SetActionBatchID(actionBatchID string) {
	o.ActionBatchID = actionBatchID
}

// WithOrganizationID adds the organizationID to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) WithOrganizationID(organizationID string) *UpdateOrganizationActionBatchParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithUpdateOrganizationActionBatch adds the updateOrganizationActionBatch to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) WithUpdateOrganizationActionBatch(updateOrganizationActionBatch UpdateOrganizationActionBatchBody) *UpdateOrganizationActionBatchParams {
	o.SetUpdateOrganizationActionBatch(updateOrganizationActionBatch)
	return o
}

// SetUpdateOrganizationActionBatch adds the updateOrganizationActionBatch to the update organization action batch params
func (o *UpdateOrganizationActionBatchParams) SetUpdateOrganizationActionBatch(updateOrganizationActionBatch UpdateOrganizationActionBatchBody) {
	o.UpdateOrganizationActionBatch = updateOrganizationActionBatch
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationActionBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionBatchId
	if err := r.SetPathParam("actionBatchId", o.ActionBatchID); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateOrganizationActionBatch); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
