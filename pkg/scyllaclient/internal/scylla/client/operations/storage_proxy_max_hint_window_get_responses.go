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

// StorageProxyMaxHintWindowGetReader is a Reader for the StorageProxyMaxHintWindowGet structure.
type StorageProxyMaxHintWindowGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMaxHintWindowGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMaxHintWindowGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMaxHintWindowGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMaxHintWindowGetOK creates a StorageProxyMaxHintWindowGetOK with default headers values
func NewStorageProxyMaxHintWindowGetOK() *StorageProxyMaxHintWindowGetOK {
	return &StorageProxyMaxHintWindowGetOK{}
}

/*StorageProxyMaxHintWindowGetOK handles this case with default header values.

StorageProxyMaxHintWindowGetOK storage proxy max hint window get o k
*/
type StorageProxyMaxHintWindowGetOK struct {
	Payload int32
}

func (o *StorageProxyMaxHintWindowGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageProxyMaxHintWindowGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMaxHintWindowGetDefault creates a StorageProxyMaxHintWindowGetDefault with default headers values
func NewStorageProxyMaxHintWindowGetDefault(code int) *StorageProxyMaxHintWindowGetDefault {
	return &StorageProxyMaxHintWindowGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyMaxHintWindowGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMaxHintWindowGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy max hint window get default response
func (o *StorageProxyMaxHintWindowGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMaxHintWindowGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMaxHintWindowGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMaxHintWindowGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}