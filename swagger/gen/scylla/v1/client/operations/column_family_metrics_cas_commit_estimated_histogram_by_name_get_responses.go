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
		result := NewColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
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

func (o *ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault creates a ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault with default headers values
func NewColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault(code int) *ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault {
	return &ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics cas commit estimated histogram by name get default response
func (o *ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsCasCommitEstimatedHistogramByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}