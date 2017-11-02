package callback

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

	"github.com/manifoldco/grafton/generated/connector/models"
)

// NewPutCallbacksIDParams creates a new PutCallbacksIDParams object
// with the default values initialized.
func NewPutCallbacksIDParams() *PutCallbacksIDParams {
	var ()
	return &PutCallbacksIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCallbacksIDParamsWithTimeout creates a new PutCallbacksIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCallbacksIDParamsWithTimeout(timeout time.Duration) *PutCallbacksIDParams {
	var ()
	return &PutCallbacksIDParams{

		timeout: timeout,
	}
}

// NewPutCallbacksIDParamsWithContext creates a new PutCallbacksIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCallbacksIDParamsWithContext(ctx context.Context) *PutCallbacksIDParams {
	var ()
	return &PutCallbacksIDParams{

		Context: ctx,
	}
}

// NewPutCallbacksIDParamsWithHTTPClient creates a new PutCallbacksIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCallbacksIDParamsWithHTTPClient(client *http.Client) *PutCallbacksIDParams {
	var ()
	return &PutCallbacksIDParams{
		HTTPClient: client,
	}
}

/*PutCallbacksIDParams contains all the parameters to send to the API endpoint
for the put callbacks ID operation typically these are written to a http.Request
*/
type PutCallbacksIDParams struct {

	/*Body
	  Response to Callback sent by the Provider

	*/
	Body *models.CallbackResponse
	/*ID
	  ID of a Callback, stored as a base32 encoded 18 byte identifier.


	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put callbacks ID params
func (o *PutCallbacksIDParams) WithTimeout(timeout time.Duration) *PutCallbacksIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put callbacks ID params
func (o *PutCallbacksIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put callbacks ID params
func (o *PutCallbacksIDParams) WithContext(ctx context.Context) *PutCallbacksIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put callbacks ID params
func (o *PutCallbacksIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put callbacks ID params
func (o *PutCallbacksIDParams) WithHTTPClient(client *http.Client) *PutCallbacksIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put callbacks ID params
func (o *PutCallbacksIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put callbacks ID params
func (o *PutCallbacksIDParams) WithBody(body *models.CallbackResponse) *PutCallbacksIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put callbacks ID params
func (o *PutCallbacksIDParams) SetBody(body *models.CallbackResponse) {
	o.Body = body
}

// WithID adds the id to the put callbacks ID params
func (o *PutCallbacksIDParams) WithID(id string) *PutCallbacksIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put callbacks ID params
func (o *PutCallbacksIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutCallbacksIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.CallbackResponse)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
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