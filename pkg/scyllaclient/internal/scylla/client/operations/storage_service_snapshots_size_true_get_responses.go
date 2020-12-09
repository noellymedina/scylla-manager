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

// StorageServiceSnapshotsSizeTrueGetReader is a Reader for the StorageServiceSnapshotsSizeTrueGet structure.
type StorageServiceSnapshotsSizeTrueGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceSnapshotsSizeTrueGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceSnapshotsSizeTrueGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceSnapshotsSizeTrueGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceSnapshotsSizeTrueGetOK creates a StorageServiceSnapshotsSizeTrueGetOK with default headers values
func NewStorageServiceSnapshotsSizeTrueGetOK() *StorageServiceSnapshotsSizeTrueGetOK {
	return &StorageServiceSnapshotsSizeTrueGetOK{}
}

/*StorageServiceSnapshotsSizeTrueGetOK handles this case with default header values.

StorageServiceSnapshotsSizeTrueGetOK storage service snapshots size true get o k
*/
type StorageServiceSnapshotsSizeTrueGetOK struct {
	Payload interface{}
}

func (o *StorageServiceSnapshotsSizeTrueGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageServiceSnapshotsSizeTrueGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceSnapshotsSizeTrueGetDefault creates a StorageServiceSnapshotsSizeTrueGetDefault with default headers values
func NewStorageServiceSnapshotsSizeTrueGetDefault(code int) *StorageServiceSnapshotsSizeTrueGetDefault {
	return &StorageServiceSnapshotsSizeTrueGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceSnapshotsSizeTrueGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceSnapshotsSizeTrueGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service snapshots size true get default response
func (o *StorageServiceSnapshotsSizeTrueGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceSnapshotsSizeTrueGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceSnapshotsSizeTrueGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceSnapshotsSizeTrueGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}