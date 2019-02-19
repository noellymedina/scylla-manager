// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGet structure.
type ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK creates a ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK() *ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK {
	return &ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK column family metrics cas commit estimated histogram by name get o k
*/
type ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/cas_commit/estimated_histogram/{name}][%d] columnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK ", 200)
}

func (o *ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
