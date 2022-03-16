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

// NewGetOrganizationSamlIdpsParams creates a new GetOrganizationSamlIdpsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationSamlIdpsParams() *GetOrganizationSamlIdpsParams {
	return &GetOrganizationSamlIdpsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationSamlIdpsParamsWithTimeout creates a new GetOrganizationSamlIdpsParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationSamlIdpsParamsWithTimeout(timeout time.Duration) *GetOrganizationSamlIdpsParams {
	return &GetOrganizationSamlIdpsParams{
		timeout: timeout,
	}
}

// NewGetOrganizationSamlIdpsParamsWithContext creates a new GetOrganizationSamlIdpsParams object
// with the ability to set a context for a request.
func NewGetOrganizationSamlIdpsParamsWithContext(ctx context.Context) *GetOrganizationSamlIdpsParams {
	return &GetOrganizationSamlIdpsParams{
		Context: ctx,
	}
}

// NewGetOrganizationSamlIdpsParamsWithHTTPClient creates a new GetOrganizationSamlIdpsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationSamlIdpsParamsWithHTTPClient(client *http.Client) *GetOrganizationSamlIdpsParams {
	return &GetOrganizationSamlIdpsParams{
		HTTPClient: client,
	}
}

/* GetOrganizationSamlIdpsParams contains all the parameters to send to the API endpoint
   for the get organization saml idps operation.

   Typically these are written to a http.Request.
*/
type GetOrganizationSamlIdpsParams struct {

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization saml idps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationSamlIdpsParams) WithDefaults() *GetOrganizationSamlIdpsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization saml idps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationSamlIdpsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization saml idps params
func (o *GetOrganizationSamlIdpsParams) WithTimeout(timeout time.Duration) *GetOrganizationSamlIdpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization saml idps params
func (o *GetOrganizationSamlIdpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization saml idps params
func (o *GetOrganizationSamlIdpsParams) WithContext(ctx context.Context) *GetOrganizationSamlIdpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization saml idps params
func (o *GetOrganizationSamlIdpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization saml idps params
func (o *GetOrganizationSamlIdpsParams) WithHTTPClient(client *http.Client) *GetOrganizationSamlIdpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization saml idps params
func (o *GetOrganizationSamlIdpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization saml idps params
func (o *GetOrganizationSamlIdpsParams) WithOrganizationID(organizationID string) *GetOrganizationSamlIdpsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization saml idps params
func (o *GetOrganizationSamlIdpsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationSamlIdpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}