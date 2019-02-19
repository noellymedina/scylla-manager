// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceRelocalSchemaPostReader is a Reader for the StorageServiceRelocalSchemaPost structure.
type StorageServiceRelocalSchemaPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRelocalSchemaPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageServiceRelocalSchemaPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceRelocalSchemaPostOK creates a StorageServiceRelocalSchemaPostOK with default headers values
func NewStorageServiceRelocalSchemaPostOK() *StorageServiceRelocalSchemaPostOK {
	return &StorageServiceRelocalSchemaPostOK{}
}

/*StorageServiceRelocalSchemaPostOK handles this case with default header values.

StorageServiceRelocalSchemaPostOK storage service relocal schema post o k
*/
type StorageServiceRelocalSchemaPostOK struct {
}

func (o *StorageServiceRelocalSchemaPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/relocal_schema][%d] storageServiceRelocalSchemaPostOK ", 200)
}

func (o *StorageServiceRelocalSchemaPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
