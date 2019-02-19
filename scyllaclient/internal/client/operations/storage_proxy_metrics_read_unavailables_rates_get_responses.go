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

// StorageProxyMetricsReadUnavailablesRatesGetReader is a Reader for the StorageProxyMetricsReadUnavailablesRatesGet structure.
type StorageProxyMetricsReadUnavailablesRatesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsReadUnavailablesRatesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageProxyMetricsReadUnavailablesRatesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyMetricsReadUnavailablesRatesGetOK creates a StorageProxyMetricsReadUnavailablesRatesGetOK with default headers values
func NewStorageProxyMetricsReadUnavailablesRatesGetOK() *StorageProxyMetricsReadUnavailablesRatesGetOK {
	return &StorageProxyMetricsReadUnavailablesRatesGetOK{}
}

/*StorageProxyMetricsReadUnavailablesRatesGetOK handles this case with default header values.

StorageProxyMetricsReadUnavailablesRatesGetOK storage proxy metrics read unavailables rates get o k
*/
type StorageProxyMetricsReadUnavailablesRatesGetOK struct {
	Payload *models.RateMovingAverage
}

func (o *StorageProxyMetricsReadUnavailablesRatesGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/metrics/read/unavailables_rates][%d] storageProxyMetricsReadUnavailablesRatesGetOK  %+v", 200, o.Payload)
}

func (o *StorageProxyMetricsReadUnavailablesRatesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RateMovingAverage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
