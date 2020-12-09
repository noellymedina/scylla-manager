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

// StorageProxyHintedHandoffEnabledGetReader is a Reader for the StorageProxyHintedHandoffEnabledGet structure.
type StorageProxyHintedHandoffEnabledGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyHintedHandoffEnabledGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyHintedHandoffEnabledGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyHintedHandoffEnabledGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyHintedHandoffEnabledGetOK creates a StorageProxyHintedHandoffEnabledGetOK with default headers values
func NewStorageProxyHintedHandoffEnabledGetOK() *StorageProxyHintedHandoffEnabledGetOK {
	return &StorageProxyHintedHandoffEnabledGetOK{}
}

/*StorageProxyHintedHandoffEnabledGetOK handles this case with default header values.

StorageProxyHintedHandoffEnabledGetOK storage proxy hinted handoff enabled get o k
*/
type StorageProxyHintedHandoffEnabledGetOK struct {
	Payload bool
}

func (o *StorageProxyHintedHandoffEnabledGetOK) GetPayload() bool {
	return o.Payload
}

func (o *StorageProxyHintedHandoffEnabledGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyHintedHandoffEnabledGetDefault creates a StorageProxyHintedHandoffEnabledGetDefault with default headers values
func NewStorageProxyHintedHandoffEnabledGetDefault(code int) *StorageProxyHintedHandoffEnabledGetDefault {
	return &StorageProxyHintedHandoffEnabledGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyHintedHandoffEnabledGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyHintedHandoffEnabledGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy hinted handoff enabled get default response
func (o *StorageProxyHintedHandoffEnabledGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyHintedHandoffEnabledGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyHintedHandoffEnabledGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyHintedHandoffEnabledGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}