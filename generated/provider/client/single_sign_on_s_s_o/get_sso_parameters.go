package single_sign_on_s_s_o

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSsoParams creates a new GetSsoParams object
// with the default values initialized.
func NewGetSsoParams() *GetSsoParams {
	var ()
	return &GetSsoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSsoParamsWithTimeout creates a new GetSsoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSsoParamsWithTimeout(timeout time.Duration) *GetSsoParams {
	var ()
	return &GetSsoParams{

		timeout: timeout,
	}
}

// NewGetSsoParamsWithContext creates a new GetSsoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSsoParamsWithContext(ctx context.Context) *GetSsoParams {
	var ()
	return &GetSsoParams{

		Context: ctx,
	}
}

// NewGetSsoParamsWithHTTPClient creates a new GetSsoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSsoParamsWithHTTPClient(client *http.Client) *GetSsoParams {
	var ()
	return &GetSsoParams{
		HTTPClient: client,
	}
}

/*GetSsoParams contains all the parameters to send to the API endpoint
for the get sso operation typically these are written to a http.Request
*/
type GetSsoParams struct {

	/*Code
	  The authorization code sent along with the SSO request which a
	provider uses to create an access token.


	*/
	Code string
	/*ResourceID
	  ID of the Resource the user is attempting to access. This ID is a
	base32 encoded 18 byte identifier.


	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sso params
func (o *GetSsoParams) WithTimeout(timeout time.Duration) *GetSsoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sso params
func (o *GetSsoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sso params
func (o *GetSsoParams) WithContext(ctx context.Context) *GetSsoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sso params
func (o *GetSsoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sso params
func (o *GetSsoParams) WithHTTPClient(client *http.Client) *GetSsoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sso params
func (o *GetSsoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCode adds the code to the get sso params
func (o *GetSsoParams) WithCode(code string) *GetSsoParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the get sso params
func (o *GetSsoParams) SetCode(code string) {
	o.Code = code
}

// WithResourceID adds the resourceID to the get sso params
func (o *GetSsoParams) WithResourceID(resourceID string) *GetSsoParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get sso params
func (o *GetSsoParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSsoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param code
	qrCode := o.Code
	qCode := qrCode
	if qCode != "" {
		if err := r.SetQueryParam("code", qCode); err != nil {
			return err
		}
	}

	// query param resource_id
	qrResourceID := o.ResourceID
	qResourceID := qrResourceID
	if qResourceID != "" {
		if err := r.SetQueryParam("resource_id", qResourceID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
