// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/models"
)

// GetRangeMetricsLatencyHistogramDepricatedReader is a Reader for the GetRangeMetricsLatencyHistogramDepricated structure.
type GetRangeMetricsLatencyHistogramDepricatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRangeMetricsLatencyHistogramDepricatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRangeMetricsLatencyHistogramDepricatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRangeMetricsLatencyHistogramDepricatedOK creates a GetRangeMetricsLatencyHistogramDepricatedOK with default headers values
func NewGetRangeMetricsLatencyHistogramDepricatedOK() *GetRangeMetricsLatencyHistogramDepricatedOK {
	return &GetRangeMetricsLatencyHistogramDepricatedOK{}
}

/*GetRangeMetricsLatencyHistogramDepricatedOK handles this case with default header values.

GetRangeMetricsLatencyHistogramDepricatedOK get range metrics latency histogram depricated o k
*/
type GetRangeMetricsLatencyHistogramDepricatedOK struct {
	Payload *models.Histogram
}

func (o *GetRangeMetricsLatencyHistogramDepricatedOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/metrics/range/histogram][%d] getRangeMetricsLatencyHistogramDepricatedOK  %+v", 200, o.Payload)
}

func (o *GetRangeMetricsLatencyHistogramDepricatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Histogram)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
