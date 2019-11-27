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

// ColumnFamilyMetricsMaxRowSizeGetReader is a Reader for the ColumnFamilyMetricsMaxRowSizeGet structure.
type ColumnFamilyMetricsMaxRowSizeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMaxRowSizeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMaxRowSizeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsMaxRowSizeGetOK creates a ColumnFamilyMetricsMaxRowSizeGetOK with default headers values
func NewColumnFamilyMetricsMaxRowSizeGetOK() *ColumnFamilyMetricsMaxRowSizeGetOK {
	return &ColumnFamilyMetricsMaxRowSizeGetOK{}
}

/*ColumnFamilyMetricsMaxRowSizeGetOK handles this case with default header values.

ColumnFamilyMetricsMaxRowSizeGetOK column family metrics max row size get o k
*/
type ColumnFamilyMetricsMaxRowSizeGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMaxRowSizeGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/max_row_size][%d] columnFamilyMetricsMaxRowSizeGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsMaxRowSizeGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMaxRowSizeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}