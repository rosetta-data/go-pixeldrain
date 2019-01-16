// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jkawamoto/go-pixeldrain/models"
)

// DeleteFileReader is a Reader for the DeleteFile structure.
type DeleteFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFileOK creates a DeleteFileOK with default headers values
func NewDeleteFileOK() *DeleteFileOK {
	return &DeleteFileOK{}
}

/*DeleteFileOK handles this case with default header values.

OK
*/
type DeleteFileOK struct {
	Payload *DeleteFileOKBody
}

func (o *DeleteFileOK) Error() string {
	return fmt.Sprintf("[DELETE /file/{id}][%d] deleteFileOK  %+v", 200, o.Payload)
}

func (o *DeleteFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteFileOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFileUnauthorized creates a DeleteFileUnauthorized with default headers values
func NewDeleteFileUnauthorized() *DeleteFileUnauthorized {
	return &DeleteFileUnauthorized{}
}

/*DeleteFileUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteFileUnauthorized struct {
	Payload *models.StandardError
}

func (o *DeleteFileUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /file/{id}][%d] deleteFileUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFileForbidden creates a DeleteFileForbidden with default headers values
func NewDeleteFileForbidden() *DeleteFileForbidden {
	return &DeleteFileForbidden{}
}

/*DeleteFileForbidden handles this case with default header values.

Forbidden
*/
type DeleteFileForbidden struct {
	Payload *models.StandardError
}

func (o *DeleteFileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /file/{id}][%d] deleteFileForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFileNotFound creates a DeleteFileNotFound with default headers values
func NewDeleteFileNotFound() *DeleteFileNotFound {
	return &DeleteFileNotFound{}
}

/*DeleteFileNotFound handles this case with default header values.

Not Found
*/
type DeleteFileNotFound struct {
	Payload *models.StandardError
}

func (o *DeleteFileNotFound) Error() string {
	return fmt.Sprintf("[DELETE /file/{id}][%d] deleteFileNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteFileOKBody delete file o k body
swagger:model DeleteFileOKBody
*/
type DeleteFileOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this delete file o k body
func (o *DeleteFileOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteFileOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteFileOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteFileOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}