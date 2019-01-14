// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFileIDDownloadParams creates a new GetFileIDDownloadParams object
// with the default values initialized.
func NewGetFileIDDownloadParams() *GetFileIDDownloadParams {
	var ()
	return &GetFileIDDownloadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFileIDDownloadParamsWithTimeout creates a new GetFileIDDownloadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFileIDDownloadParamsWithTimeout(timeout time.Duration) *GetFileIDDownloadParams {
	var ()
	return &GetFileIDDownloadParams{

		timeout: timeout,
	}
}

// NewGetFileIDDownloadParamsWithContext creates a new GetFileIDDownloadParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFileIDDownloadParamsWithContext(ctx context.Context) *GetFileIDDownloadParams {
	var ()
	return &GetFileIDDownloadParams{

		Context: ctx,
	}
}

// NewGetFileIDDownloadParamsWithHTTPClient creates a new GetFileIDDownloadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFileIDDownloadParamsWithHTTPClient(client *http.Client) *GetFileIDDownloadParams {
	var ()
	return &GetFileIDDownloadParams{
		HTTPClient: client,
	}
}

/*GetFileIDDownloadParams contains all the parameters to send to the API endpoint
for the get file ID download operation typically these are written to a http.Request
*/
type GetFileIDDownloadParams struct {

	/*ID
	  ID of the file to request

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get file ID download params
func (o *GetFileIDDownloadParams) WithTimeout(timeout time.Duration) *GetFileIDDownloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file ID download params
func (o *GetFileIDDownloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file ID download params
func (o *GetFileIDDownloadParams) WithContext(ctx context.Context) *GetFileIDDownloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file ID download params
func (o *GetFileIDDownloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file ID download params
func (o *GetFileIDDownloadParams) WithHTTPClient(client *http.Client) *GetFileIDDownloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file ID download params
func (o *GetFileIDDownloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get file ID download params
func (o *GetFileIDDownloadParams) WithID(id string) *GetFileIDDownloadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get file ID download params
func (o *GetFileIDDownloadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetFileIDDownloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
