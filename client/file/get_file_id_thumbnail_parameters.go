// Code generated by go-swagger; DO NOT EDIT.

package file

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

// NewGetFileIDThumbnailParams creates a new GetFileIDThumbnailParams object
// with the default values initialized.
func NewGetFileIDThumbnailParams() *GetFileIDThumbnailParams {
	var ()
	return &GetFileIDThumbnailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFileIDThumbnailParamsWithTimeout creates a new GetFileIDThumbnailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFileIDThumbnailParamsWithTimeout(timeout time.Duration) *GetFileIDThumbnailParams {
	var ()
	return &GetFileIDThumbnailParams{

		timeout: timeout,
	}
}

// NewGetFileIDThumbnailParamsWithContext creates a new GetFileIDThumbnailParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFileIDThumbnailParamsWithContext(ctx context.Context) *GetFileIDThumbnailParams {
	var ()
	return &GetFileIDThumbnailParams{

		Context: ctx,
	}
}

// NewGetFileIDThumbnailParamsWithHTTPClient creates a new GetFileIDThumbnailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFileIDThumbnailParamsWithHTTPClient(client *http.Client) *GetFileIDThumbnailParams {
	var ()
	return &GetFileIDThumbnailParams{
		HTTPClient: client,
	}
}

/*GetFileIDThumbnailParams contains all the parameters to send to the API endpoint
for the get file ID thumbnail operation typically these are written to a http.Request
*/
type GetFileIDThumbnailParams struct {

	/*ID
	  ID of the file to get a thumbnail for

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get file ID thumbnail params
func (o *GetFileIDThumbnailParams) WithTimeout(timeout time.Duration) *GetFileIDThumbnailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file ID thumbnail params
func (o *GetFileIDThumbnailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file ID thumbnail params
func (o *GetFileIDThumbnailParams) WithContext(ctx context.Context) *GetFileIDThumbnailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file ID thumbnail params
func (o *GetFileIDThumbnailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file ID thumbnail params
func (o *GetFileIDThumbnailParams) WithHTTPClient(client *http.Client) *GetFileIDThumbnailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file ID thumbnail params
func (o *GetFileIDThumbnailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get file ID thumbnail params
func (o *GetFileIDThumbnailParams) WithID(id string) *GetFileIDThumbnailParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get file ID thumbnail params
func (o *GetFileIDThumbnailParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetFileIDThumbnailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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