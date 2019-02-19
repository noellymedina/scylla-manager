// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceRPCServerPostReader is a Reader for the StorageServiceRPCServerPost structure.
type StorageServiceRPCServerPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRPCServerPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageServiceRPCServerPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceRPCServerPostOK creates a StorageServiceRPCServerPostOK with default headers values
func NewStorageServiceRPCServerPostOK() *StorageServiceRPCServerPostOK {
	return &StorageServiceRPCServerPostOK{}
}

/*StorageServiceRPCServerPostOK handles this case with default header values.

StorageServiceRPCServerPostOK storage service Rpc server post o k
*/
type StorageServiceRPCServerPostOK struct {
}

func (o *StorageServiceRPCServerPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/rpc_server][%d] storageServiceRpcServerPostOK ", 200)
}

func (o *StorageServiceRPCServerPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
