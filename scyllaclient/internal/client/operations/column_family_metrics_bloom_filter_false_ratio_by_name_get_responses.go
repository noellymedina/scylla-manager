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

// ColumnFamilyMetricsBloomFilterFalseRatioByNameGetReader is a Reader for the ColumnFamilyMetricsBloomFilterFalseRatioByNameGet structure.
type ColumnFamilyMetricsBloomFilterFalseRatioByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK creates a ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK with default headers values
func NewColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK() *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK {
	return &ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK{}
}

/*ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK handles this case with default header values.

ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK column family metrics bloom filter false ratio by name get o k
*/
type ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/bloom_filter_false_ratio/{name}][%d] columnFamilyMetricsBloomFilterFalseRatioByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsBloomFilterFalseRatioByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
