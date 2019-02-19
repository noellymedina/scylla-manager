// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGet structure.
type ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK creates a ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK() *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK {
	return &ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK column family metrics cas commit estimated recent histogram by name get o k
*/
type ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/cas_commit/estimated_recent_histogram/{name}][%d] columnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK ", 200)
}

func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
