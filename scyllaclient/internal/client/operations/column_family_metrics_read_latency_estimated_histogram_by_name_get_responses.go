// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGet structure.
type ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK creates a ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK() *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK {
	return &ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK column family metrics read latency estimated histogram by name get o k
*/
type ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/read_latency/estimated_histogram/{name}][%d] columnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK ", 200)
}

func (o *ColumnFamilyMetricsReadLatencyEstimatedHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
