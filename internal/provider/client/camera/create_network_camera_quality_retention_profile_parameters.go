// Code generated by go-swagger; DO NOT EDIT.

package camera

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

// NewCreateNetworkCameraQualityRetentionProfileParams creates a new CreateNetworkCameraQualityRetentionProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkCameraQualityRetentionProfileParams() *CreateNetworkCameraQualityRetentionProfileParams {
	return &CreateNetworkCameraQualityRetentionProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkCameraQualityRetentionProfileParamsWithTimeout creates a new CreateNetworkCameraQualityRetentionProfileParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkCameraQualityRetentionProfileParamsWithTimeout(timeout time.Duration) *CreateNetworkCameraQualityRetentionProfileParams {
	return &CreateNetworkCameraQualityRetentionProfileParams{
		timeout: timeout,
	}
}

// NewCreateNetworkCameraQualityRetentionProfileParamsWithContext creates a new CreateNetworkCameraQualityRetentionProfileParams object
// with the ability to set a context for a request.
func NewCreateNetworkCameraQualityRetentionProfileParamsWithContext(ctx context.Context) *CreateNetworkCameraQualityRetentionProfileParams {
	return &CreateNetworkCameraQualityRetentionProfileParams{
		Context: ctx,
	}
}

// NewCreateNetworkCameraQualityRetentionProfileParamsWithHTTPClient creates a new CreateNetworkCameraQualityRetentionProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkCameraQualityRetentionProfileParamsWithHTTPClient(client *http.Client) *CreateNetworkCameraQualityRetentionProfileParams {
	return &CreateNetworkCameraQualityRetentionProfileParams{
		HTTPClient: client,
	}
}

/* CreateNetworkCameraQualityRetentionProfileParams contains all the parameters to send to the API endpoint
   for the create network camera quality retention profile operation.

   Typically these are written to a http.Request.
*/
type CreateNetworkCameraQualityRetentionProfileParams struct {

	// CreateNetworkCameraQualityRetentionProfile.
	CreateNetworkCameraQualityRetentionProfile CreateNetworkCameraQualityRetentionProfileBody

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network camera quality retention profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkCameraQualityRetentionProfileParams) WithDefaults() *CreateNetworkCameraQualityRetentionProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network camera quality retention profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkCameraQualityRetentionProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network camera quality retention profile params
func (o *CreateNetworkCameraQualityRetentionProfileParams) WithTimeout(timeout time.Duration) *CreateNetworkCameraQualityRetentionProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network camera quality retention profile params
func (o *CreateNetworkCameraQualityRetentionProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network camera quality retention profile params
func (o *CreateNetworkCameraQualityRetentionProfileParams) WithContext(ctx context.Context) *CreateNetworkCameraQualityRetentionProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network camera quality retention profile params
func (o *CreateNetworkCameraQualityRetentionProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network camera quality retention profile params
func (o *CreateNetworkCameraQualityRetentionProfileParams) WithHTTPClient(client *http.Client) *CreateNetworkCameraQualityRetentionProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network camera quality retention profile params
func (o *CreateNetworkCameraQualityRetentionProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkCameraQualityRetentionProfile adds the createNetworkCameraQualityRetentionProfile to the create network camera quality retention profile params
func (o *CreateNetworkCameraQualityRetentionProfileParams) WithCreateNetworkCameraQualityRetentionProfile(createNetworkCameraQualityRetentionProfile CreateNetworkCameraQualityRetentionProfileBody) *CreateNetworkCameraQualityRetentionProfileParams {
	o.SetCreateNetworkCameraQualityRetentionProfile(createNetworkCameraQualityRetentionProfile)
	return o
}

// SetCreateNetworkCameraQualityRetentionProfile adds the createNetworkCameraQualityRetentionProfile to the create network camera quality retention profile params
func (o *CreateNetworkCameraQualityRetentionProfileParams) SetCreateNetworkCameraQualityRetentionProfile(createNetworkCameraQualityRetentionProfile CreateNetworkCameraQualityRetentionProfileBody) {
	o.CreateNetworkCameraQualityRetentionProfile = createNetworkCameraQualityRetentionProfile
}

// WithNetworkID adds the networkID to the create network camera quality retention profile params
func (o *CreateNetworkCameraQualityRetentionProfileParams) WithNetworkID(networkID string) *CreateNetworkCameraQualityRetentionProfileParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network camera quality retention profile params
func (o *CreateNetworkCameraQualityRetentionProfileParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkCameraQualityRetentionProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateNetworkCameraQualityRetentionProfile); err != nil {
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
