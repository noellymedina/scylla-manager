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

// ColumnFamilyMetricsMemtableOnHeapSizeByNameGetReader is a Reader for the ColumnFamilyMetricsMemtableOnHeapSizeByNameGet structure.
type ColumnFamilyMetricsMemtableOnHeapSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK creates a ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK() *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK {
	return &ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK{}
}

/*ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK column family metrics memtable on heap size by name get o k
*/
type ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/memtable_on_heap_size/{name}][%d] columnFamilyMetricsMemtableOnHeapSizeByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsMemtableOnHeapSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
