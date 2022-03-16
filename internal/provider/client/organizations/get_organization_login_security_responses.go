// Code generated by go-swagger; DO NOT EDIT.

package organizations

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
)

// GetOrganizationLoginSecurityReader is a Reader for the GetOrganizationLoginSecurity structure.
type GetOrganizationLoginSecurityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationLoginSecurityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationLoginSecurityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationLoginSecurityOK creates a GetOrganizationLoginSecurityOK with default headers values
func NewGetOrganizationLoginSecurityOK() *GetOrganizationLoginSecurityOK {
	return &GetOrganizationLoginSecurityOK{}
}

/* GetOrganizationLoginSecurityOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationLoginSecurityOK struct {
	Payload *GetOrganizationLoginSecurityOKBody
}

func (o *GetOrganizationLoginSecurityOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/loginSecurity][%d] getOrganizationLoginSecurityOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationLoginSecurityOK) GetPayload() *GetOrganizationLoginSecurityOKBody {
	return o.Payload
}

func (o *GetOrganizationLoginSecurityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrganizationLoginSecurityOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrganizationLoginSecurityOKBody get organization login security o k body
swagger:model GetOrganizationLoginSecurityOKBody
*/
type GetOrganizationLoginSecurityOKBody struct {

	// Number of failed login attempts before the admin is locked out of the org
	AccountLockoutAttempts int64 `json:"accountLockoutAttempts,omitempty"`

	// api authentication
	APIAuthentication *GetOrganizationLoginSecurityOKBodyAPIAuthentication `json:"apiAuthentication,omitempty"`

	// Whether account lock-out is enabled for the org
	EnforceAccountLockout bool `json:"enforceAccountLockout,omitempty"`

	// Whether the org is forced to choose a password that is distinct from previous passwords
	EnforceDifferentPasswords bool `json:"enforceDifferentPasswords,omitempty"`

	// Whether the org admin should be logged out after idling for the given idle timeout duration
	EnforceIdleTimeout bool `json:"enforceIdleTimeout,omitempty"`

	// Whether the org restricts access to Dashboard by IP addresses
	EnforceLoginIPRanges bool `json:"enforceLoginIpRanges,omitempty"`

	// Whether password expiration should be enforced
	EnforcePasswordExpiration bool `json:"enforcePasswordExpiration,omitempty"`

	// Whether the org must choose a strong password
	EnforceStrongPasswords bool `json:"enforceStrongPasswords,omitempty"`

	// Whether the org uses two factor authentication
	EnforceTwoFactorAuth bool `json:"enforceTwoFactorAuth,omitempty"`

	// Length of time, in minutes, an org admin is idle before being logged out
	IdleTimeoutMinutes int64 `json:"idleTimeoutMinutes,omitempty"`

	// The list of IP addresses that are allowed to access to Dashboard
	LoginIPRanges []string `json:"loginIpRanges"`

	// How many recent password that cannot be reused
	NumDifferentPasswords int64 `json:"numDifferentPasswords,omitempty"`

	// How long the password is valid, in days
	PasswordExpirationDays int64 `json:"passwordExpirationDays,omitempty"`
}

// Validate validates this get organization login security o k body
func (o *GetOrganizationLoginSecurityOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAPIAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationLoginSecurityOKBody) validateAPIAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(o.APIAuthentication) { // not required
		return nil
	}

	if o.APIAuthentication != nil {
		if err := o.APIAuthentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationLoginSecurityOK" + "." + "apiAuthentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationLoginSecurityOK" + "." + "apiAuthentication")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get organization login security o k body based on the context it is used
func (o *GetOrganizationLoginSecurityOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAPIAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationLoginSecurityOKBody) contextValidateAPIAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if o.APIAuthentication != nil {
		if err := o.APIAuthentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationLoginSecurityOK" + "." + "apiAuthentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationLoginSecurityOK" + "." + "apiAuthentication")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationLoginSecurityOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationLoginSecurityOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganizationLoginSecurityOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOrganizationLoginSecurityOKBodyAPIAuthentication Collection of properties related to API authentication
swagger:model GetOrganizationLoginSecurityOKBodyAPIAuthentication
*/
type GetOrganizationLoginSecurityOKBodyAPIAuthentication struct {

	// ip restrictions for keys
	IPRestrictionsForKeys *GetOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"`
}

// Validate validates this get organization login security o k body API authentication
func (o *GetOrganizationLoginSecurityOKBodyAPIAuthentication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPRestrictionsForKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationLoginSecurityOKBodyAPIAuthentication) validateIPRestrictionsForKeys(formats strfmt.Registry) error {
	if swag.IsZero(o.IPRestrictionsForKeys) { // not required
		return nil
	}

	if o.IPRestrictionsForKeys != nil {
		if err := o.IPRestrictionsForKeys.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationLoginSecurityOK" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationLoginSecurityOK" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get organization login security o k body API authentication based on the context it is used
func (o *GetOrganizationLoginSecurityOKBodyAPIAuthentication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPRestrictionsForKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationLoginSecurityOKBodyAPIAuthentication) contextValidateIPRestrictionsForKeys(ctx context.Context, formats strfmt.Registry) error {

	if o.IPRestrictionsForKeys != nil {
		if err := o.IPRestrictionsForKeys.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationLoginSecurityOK" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationLoginSecurityOK" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationLoginSecurityOKBodyAPIAuthentication) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationLoginSecurityOKBodyAPIAuthentication) UnmarshalBinary(b []byte) error {
	var res GetOrganizationLoginSecurityOKBodyAPIAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys Collection of restrictions by IP address
swagger:model GetOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys
*/
type GetOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys struct {

	// Whether API IP restrictions are enabled
	Enabled bool `json:"enabled,omitempty"`

	// The list of IP address ranges that are allowed to make API requests for the org
	Ranges []string `json:"ranges"`
}

// Validate validates this get organization login security o k body API authentication IP restrictions for keys
func (o *GetOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization login security o k body API authentication IP restrictions for keys based on context it is used
func (o *GetOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys) UnmarshalBinary(b []byte) error {
	var res GetOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}