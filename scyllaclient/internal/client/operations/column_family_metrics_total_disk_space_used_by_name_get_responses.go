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

// ColumnFamilyMetricsTotalDiskSpaceUsedByNameGetReader is a Reader for the ColumnFamilyMetricsTotalDiskSpaceUsedByNameGet structure.
type ColumnFamilyMetricsTotalDiskSpaceUsedByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsTotalDiskSpaceUsedByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsTotalDiskSpaceUsedByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsTotalDiskSpaceUsedByNameGetOK creates a ColumnFamilyMetricsTotalDiskSpaceUsedByNameGetOK with default headers values
func NewColumnFamilyMetricsTotalDiskSpaceUsedByNameGetOK() *ColumnFamilyMetricsTotalDiskSpaceUsedByNameGetOK {
	return &ColumnFamilyMetricsTotalDiskSpaceUsedByNameGetOK{}
}

/*ColumnFamilyMetricsTotalDiskSpaceUsedByNameGetOK handles this case with default header values.

ColumnFamilyMetricsTotalDiskSpaceUsedByNameGetOK column family metrics total disk space used by name get o k
*/
type ColumnFamilyMetricsTotalDiskSpaceUsedByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsTotalDiskSpaceUsedByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/total_disk_space_used/{name}][%d] columnFamilyMetricsTotalDiskSpaceUsedByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsTotalDiskSpaceUsedByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
