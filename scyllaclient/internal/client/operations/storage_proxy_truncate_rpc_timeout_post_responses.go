// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageProxyTruncateRPCTimeoutPostReader is a Reader for the StorageProxyTruncateRPCTimeoutPost structure.
type StorageProxyTruncateRPCTimeoutPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyTruncateRPCTimeoutPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageProxyTruncateRPCTimeoutPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyTruncateRPCTimeoutPostOK creates a StorageProxyTruncateRPCTimeoutPostOK with default headers values
func NewStorageProxyTruncateRPCTimeoutPostOK() *StorageProxyTruncateRPCTimeoutPostOK {
	return &StorageProxyTruncateRPCTimeoutPostOK{}
}

/*StorageProxyTruncateRPCTimeoutPostOK handles this case with default header values.

StorageProxyTruncateRPCTimeoutPostOK storage proxy truncate Rpc timeout post o k
*/
type StorageProxyTruncateRPCTimeoutPostOK struct {
}

func (o *StorageProxyTruncateRPCTimeoutPostOK) Error() string {
	return fmt.Sprintf("[POST /storage_proxy/truncate_rpc_timeout][%d] storageProxyTruncateRpcTimeoutPostOK ", 200)
}

func (o *StorageProxyTruncateRPCTimeoutPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
