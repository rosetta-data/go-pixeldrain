// Code generated by go-swagger; DO NOT EDIT.

package list

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jkawamoto/go-pixeldrain/models"
)

// PostListReader is a Reader for the PostList structure.
type PostListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 413:
		result := NewPostListRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostListUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostListOK creates a PostListOK with default headers values
func NewPostListOK() *PostListOK {
	return &PostListOK{}
}

/*PostListOK handles this case with default header values.

OK
*/
type PostListOK struct {
	Payload *models.PostListOKBody
}

func (o *PostListOK) Error() string {
	return fmt.Sprintf("[POST /list][%d] postListOK  %+v", 200, o.Payload)
}

func (o *PostListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostListRequestEntityTooLarge creates a PostListRequestEntityTooLarge with default headers values
func NewPostListRequestEntityTooLarge() *PostListRequestEntityTooLarge {
	return &PostListRequestEntityTooLarge{}
}

/*PostListRequestEntityTooLarge handles this case with default header values.

Payload too large
*/
type PostListRequestEntityTooLarge struct {
	Payload *models.StandardError
}

func (o *PostListRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /list][%d] postListRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostListRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostListUnprocessableEntity creates a PostListUnprocessableEntity with default headers values
func NewPostListUnprocessableEntity() *PostListUnprocessableEntity {
	return &PostListUnprocessableEntity{}
}

/*PostListUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PostListUnprocessableEntity struct {
	Payload *models.StandardError
}

func (o *PostListUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /list][%d] postListUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostListUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}