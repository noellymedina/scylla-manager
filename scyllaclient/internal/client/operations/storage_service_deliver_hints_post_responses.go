// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceDeliverHintsPostReader is a Reader for the StorageServiceDeliverHintsPost structure.
type StorageServiceDeliverHintsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceDeliverHintsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageServiceDeliverHintsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceDeliverHintsPostOK creates a StorageServiceDeliverHintsPostOK with default headers values
func NewStorageServiceDeliverHintsPostOK() *StorageServiceDeliverHintsPostOK {
	return &StorageServiceDeliverHintsPostOK{}
}

/*StorageServiceDeliverHintsPostOK handles this case with default header values.

StorageServiceDeliverHintsPostOK storage service deliver hints post o k
*/
type StorageServiceDeliverHintsPostOK struct {
}

func (o *StorageServiceDeliverHintsPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/deliver_hints][%d] storageServiceDeliverHintsPostOK ", 200)
}

func (o *StorageServiceDeliverHintsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
