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

// StorageProxyMetricsCasReadTimeoutsGetReader is a Reader for the StorageProxyMetricsCasReadTimeoutsGet structure.
type StorageProxyMetricsCasReadTimeoutsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsCasReadTimeoutsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageProxyMetricsCasReadTimeoutsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyMetricsCasReadTimeoutsGetOK creates a StorageProxyMetricsCasReadTimeoutsGetOK with default headers values
func NewStorageProxyMetricsCasReadTimeoutsGetOK() *StorageProxyMetricsCasReadTimeoutsGetOK {
	return &StorageProxyMetricsCasReadTimeoutsGetOK{}
}

/*StorageProxyMetricsCasReadTimeoutsGetOK handles this case with default header values.

StorageProxyMetricsCasReadTimeoutsGetOK storage proxy metrics cas read timeouts get o k
*/
type StorageProxyMetricsCasReadTimeoutsGetOK struct {
	Payload interface{}
}

func (o *StorageProxyMetricsCasReadTimeoutsGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/metrics/cas_read/timeouts][%d] storageProxyMetricsCasReadTimeoutsGetOK  %+v", 200, o.Payload)
}

func (o *StorageProxyMetricsCasReadTimeoutsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
