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

// ColumnFamilyMetricsMemtableColumnsCountGetReader is a Reader for the ColumnFamilyMetricsMemtableColumnsCountGet structure.
type ColumnFamilyMetricsMemtableColumnsCountGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMemtableColumnsCountGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsMemtableColumnsCountGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsMemtableColumnsCountGetOK creates a ColumnFamilyMetricsMemtableColumnsCountGetOK with default headers values
func NewColumnFamilyMetricsMemtableColumnsCountGetOK() *ColumnFamilyMetricsMemtableColumnsCountGetOK {
	return &ColumnFamilyMetricsMemtableColumnsCountGetOK{}
}

/*ColumnFamilyMetricsMemtableColumnsCountGetOK handles this case with default header values.

ColumnFamilyMetricsMemtableColumnsCountGetOK column family metrics memtable columns count get o k
*/
type ColumnFamilyMetricsMemtableColumnsCountGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMemtableColumnsCountGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/memtable_columns_count/][%d] columnFamilyMetricsMemtableColumnsCountGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsMemtableColumnsCountGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
