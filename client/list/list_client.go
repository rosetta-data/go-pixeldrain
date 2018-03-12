// Code generated by go-swagger; DO NOT EDIT.

package list

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new list API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for list API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetListID returns information about a file list and the files in it

Returns information about a file list and the files in it.
*/
func (a *Client) GetListID(params *GetListIDParams) (*GetListIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetListIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetListID",
		Method:             "GET",
		PathPattern:        "/list/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "test/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetListIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetListIDOK), nil

}

/*
PostList creates a list of files that can be viewed together on the file viewer page

Creates a list of files that can be viewed together on the file viewer page.
*/
func (a *Client) PostList(params *PostListParams) (*PostListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostList",
		Method:             "POST",
		PathPattern:        "/list",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "test/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}