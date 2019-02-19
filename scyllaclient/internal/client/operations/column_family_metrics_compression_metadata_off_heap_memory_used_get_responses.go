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

// ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetReader is a Reader for the ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGet structure.
type ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK creates a ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK with default headers values
func NewColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK() *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK {
	return &ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK{}
}

/*ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK handles this case with default header values.

ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK column family metrics compression metadata off heap memory used get o k
*/
type ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/compression_metadata_off_heap_memory_used][%d] columnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsCompressionMetadataOffHeapMemoryUsedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
