// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageProxyReadRepairAttemptedGetReader is a Reader for the StorageProxyReadRepairAttemptedGet structure.
type StorageProxyReadRepairAttemptedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyReadRepairAttemptedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageProxyReadRepairAttemptedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyReadRepairAttemptedGetOK creates a StorageProxyReadRepairAttemptedGetOK with default headers values
func NewStorageProxyReadRepairAttemptedGetOK() *StorageProxyReadRepairAttemptedGetOK {
	return &StorageProxyReadRepairAttemptedGetOK{}
}

/*StorageProxyReadRepairAttemptedGetOK handles this case with default header values.

StorageProxyReadRepairAttemptedGetOK storage proxy read repair attempted get o k
*/
type StorageProxyReadRepairAttemptedGetOK struct {
	Payload interface{}
}

func (o *StorageProxyReadRepairAttemptedGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/read_repair_attempted][%d] storageProxyReadRepairAttemptedGetOK  %+v", 200, o.Payload)
}

func (o *StorageProxyReadRepairAttemptedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
