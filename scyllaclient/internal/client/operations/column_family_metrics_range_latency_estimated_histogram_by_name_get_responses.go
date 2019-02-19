// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGet structure.
type ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK creates a ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK() *ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK {
	return &ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK column family metrics range latency estimated histogram by name get o k
*/
type ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/range_latency/estimated_histogram/{name}][%d] columnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK ", 200)
}

func (o *ColumnFamilyMetricsRangeLatencyEstimatedHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
