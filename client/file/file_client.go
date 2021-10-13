// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new file API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for file API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteFile(params *DeleteFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFileOK, error)

	DownloadFile(params *DownloadFileParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*DownloadFileOK, error)

	GetFileInfo(params *GetFileInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFileInfoOK, error)

	GetFileThumbnail(params *GetFileThumbnailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFileThumbnailOK, error)

	UploadFile(params *UploadFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadFileCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteFile deletes a file

  Deletes a file. Only works when the users owns the file.
*/
func (a *Client) DeleteFile(params *DeleteFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteFile",
		Method:             "DELETE",
		PathPattern:        "/file/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DownloadFile downloads a file

  Returns the full file associated with the ID. Supports byte range requests.
Warning: If a file is using too much bandwidth it can be rate limited. The rate limit will be enabled if a file has three times more downloads than views. The owner of a file can always download it. When a file is rate limited the user will need to fill out a captcha in order to continue downloading the file. The captcha will only appear on the file viewer page (pixeldrain.com/u/{id}). Rate limiting has been added to prevent the spread of viruses and to stop hotlinking. Hotlinking is only allowed when files are uploaded using a Pro account.
Pixeldrain also includes a virus scanner. If a virus has been detected in a file the user will also have to fill in a captcha to download it.

*/
func (a *Client) DownloadFile(params *DownloadFileParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*DownloadFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadFile",
		Method:             "GET",
		PathPattern:        "/file/{id}",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadFileReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DownloadFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DownloadFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFileInfo retrieves information of a file

  Returns information about one or more files. You can also put a comma separated list of file IDs in the URL and it will return an array of file info, instead of a single object.

*/
func (a *Client) GetFileInfo(params *GetFileInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFileInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFileInfo",
		Method:             "GET",
		PathPattern:        "/file/{id}/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFileInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFileInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFileInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFileThumbnail gets a thumbnail image representing the file

  Returns a PNG thumbnail image representing the file. The thumbnail is always 100*100 px. If the source file is parsable by imagemagick the thumbnail will be generated from the file, if not it will be a generic mime type icon.

*/
func (a *Client) GetFileThumbnail(params *GetFileThumbnailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFileThumbnailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileThumbnailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFileThumbnail",
		Method:             "GET",
		PathPattern:        "/file/{id}/thumbnail",
		ProducesMediaTypes: []string{"image/png"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFileThumbnailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFileThumbnailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFileThumbnailDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UploadFile uploads a file

  Upload a file. I recommend that you use the PUT API instead of the POST API. It’s easier to use and the multipart encoding of the POST API can cause performance issues in certain environments.

*/
func (a *Client) UploadFile(params *UploadFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadFileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadFile",
		Method:             "POST",
		PathPattern:        "/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadFileCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UploadFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
