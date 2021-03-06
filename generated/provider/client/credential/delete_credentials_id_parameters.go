package credential

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

// NewDeleteCredentialsIDParams creates a new DeleteCredentialsIDParams object
// with the default values initialized.
func NewDeleteCredentialsIDParams() *DeleteCredentialsIDParams {
	var ()
	return &DeleteCredentialsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCredentialsIDParamsWithTimeout creates a new DeleteCredentialsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCredentialsIDParamsWithTimeout(timeout time.Duration) *DeleteCredentialsIDParams {
	var ()
	return &DeleteCredentialsIDParams{

		timeout: timeout,
	}
}

// NewDeleteCredentialsIDParamsWithContext creates a new DeleteCredentialsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCredentialsIDParamsWithContext(ctx context.Context) *DeleteCredentialsIDParams {
	var ()
	return &DeleteCredentialsIDParams{

		Context: ctx,
	}
}

// NewDeleteCredentialsIDParamsWithHTTPClient creates a new DeleteCredentialsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCredentialsIDParamsWithHTTPClient(client *http.Client) *DeleteCredentialsIDParams {
	var ()
	return &DeleteCredentialsIDParams{
		HTTPClient: client,
	}
}

/*DeleteCredentialsIDParams contains all the parameters to send to the API endpoint
for the delete credentials ID operation typically these are written to a http.Request
*/
type DeleteCredentialsIDParams struct {

	/*Date
	  Timestamp of when the request was issued from Manifold in UTC.

	*/
	Date strfmt.DateTime
	/*XCallbackID
	  ID of the Callback for completing this request if the provider returns a
	`202 Accepted`, stored as a base 32 encoded 18 byte identifier.


	*/
	XCallbackID string
	/*XCallbackURL
	  The URL the provider calls to complete the request if a `202 Accepted` is
	returned.


	*/
	XCallbackURL string
	/*XSignature
	  Header containing the signature, ephemeral public key, and
	signature of the used public key signed by the Manifold root
	signing key to prove authenticity of the request.

	```
	X-Signature: [request-signature] [public-key-value] [signature-of-public-key]
	```

	examples:

	```
	X-Signature: 96afMc5FVZjhGQ4YLPyRW9tcYoPKyj1EMUxkzma_Q4c WydEYGQb7Y4ER6KPAc-YuIwAqFUpII5P9U3MAZ3wxAQ ozhcosOmuWltP8r1BAs-0kccV1AkbHcKYLSjU0dGUHY
	```


	*/
	XSignature string
	/*XSignedHeaders
	  Comma delimited ordered list of header fields used in generating
	the request signature.


	*/
	XSignedHeaders string
	/*ID
	  ID of a Credential object, stored as a base32 encoded 18 byte identifier.


	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete credentials ID params
func (o *DeleteCredentialsIDParams) WithTimeout(timeout time.Duration) *DeleteCredentialsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete credentials ID params
func (o *DeleteCredentialsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete credentials ID params
func (o *DeleteCredentialsIDParams) WithContext(ctx context.Context) *DeleteCredentialsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete credentials ID params
func (o *DeleteCredentialsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete credentials ID params
func (o *DeleteCredentialsIDParams) WithHTTPClient(client *http.Client) *DeleteCredentialsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete credentials ID params
func (o *DeleteCredentialsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDate adds the date to the delete credentials ID params
func (o *DeleteCredentialsIDParams) WithDate(date strfmt.DateTime) *DeleteCredentialsIDParams {
	o.SetDate(date)
	return o
}

// SetDate adds the date to the delete credentials ID params
func (o *DeleteCredentialsIDParams) SetDate(date strfmt.DateTime) {
	o.Date = date
}

// WithXCallbackID adds the xCallbackID to the delete credentials ID params
func (o *DeleteCredentialsIDParams) WithXCallbackID(xCallbackID string) *DeleteCredentialsIDParams {
	o.SetXCallbackID(xCallbackID)
	return o
}

// SetXCallbackID adds the xCallbackId to the delete credentials ID params
func (o *DeleteCredentialsIDParams) SetXCallbackID(xCallbackID string) {
	o.XCallbackID = xCallbackID
}

// WithXCallbackURL adds the xCallbackURL to the delete credentials ID params
func (o *DeleteCredentialsIDParams) WithXCallbackURL(xCallbackURL string) *DeleteCredentialsIDParams {
	o.SetXCallbackURL(xCallbackURL)
	return o
}

// SetXCallbackURL adds the xCallbackUrl to the delete credentials ID params
func (o *DeleteCredentialsIDParams) SetXCallbackURL(xCallbackURL string) {
	o.XCallbackURL = xCallbackURL
}

// WithXSignature adds the xSignature to the delete credentials ID params
func (o *DeleteCredentialsIDParams) WithXSignature(xSignature string) *DeleteCredentialsIDParams {
	o.SetXSignature(xSignature)
	return o
}

// SetXSignature adds the xSignature to the delete credentials ID params
func (o *DeleteCredentialsIDParams) SetXSignature(xSignature string) {
	o.XSignature = xSignature
}

// WithXSignedHeaders adds the xSignedHeaders to the delete credentials ID params
func (o *DeleteCredentialsIDParams) WithXSignedHeaders(xSignedHeaders string) *DeleteCredentialsIDParams {
	o.SetXSignedHeaders(xSignedHeaders)
	return o
}

// SetXSignedHeaders adds the xSignedHeaders to the delete credentials ID params
func (o *DeleteCredentialsIDParams) SetXSignedHeaders(xSignedHeaders string) {
	o.XSignedHeaders = xSignedHeaders
}

// WithID adds the id to the delete credentials ID params
func (o *DeleteCredentialsIDParams) WithID(id string) *DeleteCredentialsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete credentials ID params
func (o *DeleteCredentialsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCredentialsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Date
	if err := r.SetHeaderParam("Date", o.Date.String()); err != nil {
		return err
	}

	// header param X-Callback-ID
	if err := r.SetHeaderParam("X-Callback-ID", o.XCallbackID); err != nil {
		return err
	}

	// header param X-Callback-URL
	if err := r.SetHeaderParam("X-Callback-URL", o.XCallbackURL); err != nil {
		return err
	}

	// header param X-Signature
	if err := r.SetHeaderParam("X-Signature", o.XSignature); err != nil {
		return err
	}

	// header param X-Signed-Headers
	if err := r.SetHeaderParam("X-Signed-Headers", o.XSignedHeaders); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
