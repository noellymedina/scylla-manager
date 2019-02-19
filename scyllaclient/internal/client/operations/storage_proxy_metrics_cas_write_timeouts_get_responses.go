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

// StorageProxyMetricsCasWriteTimeoutsGetReader is a Reader for the StorageProxyMetricsCasWriteTimeoutsGet structure.
type StorageProxyMetricsCasWriteTimeoutsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsCasWriteTimeoutsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageProxyMetricsCasWriteTimeoutsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyMetricsCasWriteTimeoutsGetOK creates a StorageProxyMetricsCasWriteTimeoutsGetOK with default headers values
func NewStorageProxyMetricsCasWriteTimeoutsGetOK() *StorageProxyMetricsCasWriteTimeoutsGetOK {
	return &StorageProxyMetricsCasWriteTimeoutsGetOK{}
}

/*StorageProxyMetricsCasWriteTimeoutsGetOK handles this case with default header values.

StorageProxyMetricsCasWriteTimeoutsGetOK storage proxy metrics cas write timeouts get o k
*/
type StorageProxyMetricsCasWriteTimeoutsGetOK struct {
	Payload interface{}
}

func (o *StorageProxyMetricsCasWriteTimeoutsGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/metrics/cas_write/timeouts][%d] storageProxyMetricsCasWriteTimeoutsGetOK  %+v", 200, o.Payload)
}

func (o *StorageProxyMetricsCasWriteTimeoutsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
