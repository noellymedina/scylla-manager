// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla/models"
)

// StorageProxyMetricsRangeTimeoutsRatesGetReader is a Reader for the StorageProxyMetricsRangeTimeoutsRatesGet structure.
type StorageProxyMetricsRangeTimeoutsRatesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsRangeTimeoutsRatesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsRangeTimeoutsRatesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsRangeTimeoutsRatesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsRangeTimeoutsRatesGetOK creates a StorageProxyMetricsRangeTimeoutsRatesGetOK with default headers values
func NewStorageProxyMetricsRangeTimeoutsRatesGetOK() *StorageProxyMetricsRangeTimeoutsRatesGetOK {
	return &StorageProxyMetricsRangeTimeoutsRatesGetOK{}
}

/*StorageProxyMetricsRangeTimeoutsRatesGetOK handles this case with default header values.

StorageProxyMetricsRangeTimeoutsRatesGetOK storage proxy metrics range timeouts rates get o k
*/
type StorageProxyMetricsRangeTimeoutsRatesGetOK struct {
	Payload *models.RateMovingAverage
}

func (o *StorageProxyMetricsRangeTimeoutsRatesGetOK) GetPayload() *models.RateMovingAverage {
	return o.Payload
}

func (o *StorageProxyMetricsRangeTimeoutsRatesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RateMovingAverage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsRangeTimeoutsRatesGetDefault creates a StorageProxyMetricsRangeTimeoutsRatesGetDefault with default headers values
func NewStorageProxyMetricsRangeTimeoutsRatesGetDefault(code int) *StorageProxyMetricsRangeTimeoutsRatesGetDefault {
	return &StorageProxyMetricsRangeTimeoutsRatesGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyMetricsRangeTimeoutsRatesGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsRangeTimeoutsRatesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics range timeouts rates get default response
func (o *StorageProxyMetricsRangeTimeoutsRatesGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsRangeTimeoutsRatesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsRangeTimeoutsRatesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsRangeTimeoutsRatesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}