// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/models"
)

// StorageServiceTokensEndpointGetReader is a Reader for the StorageServiceTokensEndpointGet structure.
type StorageServiceTokensEndpointGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceTokensEndpointGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageServiceTokensEndpointGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceTokensEndpointGetOK creates a StorageServiceTokensEndpointGetOK with default headers values
func NewStorageServiceTokensEndpointGetOK() *StorageServiceTokensEndpointGetOK {
	return &StorageServiceTokensEndpointGetOK{}
}

/*StorageServiceTokensEndpointGetOK handles this case with default header values.

StorageServiceTokensEndpointGetOK storage service tokens endpoint get o k
*/
type StorageServiceTokensEndpointGetOK struct {
	Payload []*models.Mapper
}

func (o *StorageServiceTokensEndpointGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/tokens_endpoint][%d] storageServiceTokensEndpointGetOK  %+v", 200, o.Payload)
}

func (o *StorageServiceTokensEndpointGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
