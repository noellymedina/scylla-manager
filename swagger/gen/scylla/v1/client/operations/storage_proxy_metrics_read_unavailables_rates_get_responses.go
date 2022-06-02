// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v1/models"
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
		result := NewStorageProxyMetricsReadUnavailablesRatesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
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

func (o *StorageProxyMetricsReadUnavailablesRatesGetOK) GetPayload() *models.RateMovingAverage {
	return o.Payload
}

func (o *StorageProxyMetricsReadUnavailablesRatesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RateMovingAverage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsReadUnavailablesRatesGetDefault creates a StorageProxyMetricsReadUnavailablesRatesGetDefault with default headers values
func NewStorageProxyMetricsReadUnavailablesRatesGetDefault(code int) *StorageProxyMetricsReadUnavailablesRatesGetDefault {
	return &StorageProxyMetricsReadUnavailablesRatesGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyMetricsReadUnavailablesRatesGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsReadUnavailablesRatesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics read unavailables rates get default response
func (o *StorageProxyMetricsReadUnavailablesRatesGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsReadUnavailablesRatesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsReadUnavailablesRatesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsReadUnavailablesRatesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}