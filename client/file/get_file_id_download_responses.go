// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jkawamoto/go-pixeldrain/models"
)

// GetFileIDDownloadReader is a Reader for the GetFileIDDownload structure.
type GetFileIDDownloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileIDDownloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFileIDDownloadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetFileIDDownloadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFileIDDownloadOK creates a GetFileIDDownloadOK with default headers values
func NewGetFileIDDownloadOK() *GetFileIDDownloadOK {
	return &GetFileIDDownloadOK{}
}

/*GetFileIDDownloadOK handles this case with default header values.

A file output stream.
*/
type GetFileIDDownloadOK struct {
}

func (o *GetFileIDDownloadOK) Error() string {
	return fmt.Sprintf("[GET /file/{id}/download][%d] getFileIdDownloadOK ", 200)
}

func (o *GetFileIDDownloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFileIDDownloadDefault creates a GetFileIDDownloadDefault with default headers values
func NewGetFileIDDownloadDefault(code int) *GetFileIDDownloadDefault {
	return &GetFileIDDownloadDefault{
		_statusCode: code,
	}
}

/*GetFileIDDownloadDefault handles this case with default header values.

Error Response
*/
type GetFileIDDownloadDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get file ID download default response
func (o *GetFileIDDownloadDefault) Code() int {
	return o._statusCode
}

func (o *GetFileIDDownloadDefault) Error() string {
	return fmt.Sprintf("[GET /file/{id}/download][%d] GetFileIDDownload default  %+v", o._statusCode, o.Payload)
}

func (o *GetFileIDDownloadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
