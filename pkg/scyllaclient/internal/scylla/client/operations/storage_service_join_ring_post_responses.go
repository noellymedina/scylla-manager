// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla/models"
)

// StorageServiceJoinRingPostReader is a Reader for the StorageServiceJoinRingPost structure.
type StorageServiceJoinRingPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceJoinRingPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceJoinRingPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceJoinRingPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceJoinRingPostOK creates a StorageServiceJoinRingPostOK with default headers values
func NewStorageServiceJoinRingPostOK() *StorageServiceJoinRingPostOK {
	return &StorageServiceJoinRingPostOK{}
}

/*StorageServiceJoinRingPostOK handles this case with default header values.

StorageServiceJoinRingPostOK storage service join ring post o k
*/
type StorageServiceJoinRingPostOK struct {
}

func (o *StorageServiceJoinRingPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceJoinRingPostDefault creates a StorageServiceJoinRingPostDefault with default headers values
func NewStorageServiceJoinRingPostDefault(code int) *StorageServiceJoinRingPostDefault {
	return &StorageServiceJoinRingPostDefault{
		_statusCode: code,
	}
}

/*StorageServiceJoinRingPostDefault handles this case with default header values.

internal server error
*/
type StorageServiceJoinRingPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service join ring post default response
func (o *StorageServiceJoinRingPostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceJoinRingPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceJoinRingPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceJoinRingPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}