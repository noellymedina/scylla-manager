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

// ColumnFamilyMetricsCompressionRatioGetReader is a Reader for the ColumnFamilyMetricsCompressionRatioGet structure.
type ColumnFamilyMetricsCompressionRatioGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsCompressionRatioGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsCompressionRatioGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsCompressionRatioGetOK creates a ColumnFamilyMetricsCompressionRatioGetOK with default headers values
func NewColumnFamilyMetricsCompressionRatioGetOK() *ColumnFamilyMetricsCompressionRatioGetOK {
	return &ColumnFamilyMetricsCompressionRatioGetOK{}
}

/*ColumnFamilyMetricsCompressionRatioGetOK handles this case with default header values.

ColumnFamilyMetricsCompressionRatioGetOK column family metrics compression ratio get o k
*/
type ColumnFamilyMetricsCompressionRatioGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsCompressionRatioGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/compression_ratio][%d] columnFamilyMetricsCompressionRatioGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsCompressionRatioGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
