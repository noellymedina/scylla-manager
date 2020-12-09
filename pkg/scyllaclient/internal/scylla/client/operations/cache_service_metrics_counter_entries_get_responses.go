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

// CacheServiceMetricsCounterEntriesGetReader is a Reader for the CacheServiceMetricsCounterEntriesGet structure.
type CacheServiceMetricsCounterEntriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsCounterEntriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsCounterEntriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsCounterEntriesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsCounterEntriesGetOK creates a CacheServiceMetricsCounterEntriesGetOK with default headers values
func NewCacheServiceMetricsCounterEntriesGetOK() *CacheServiceMetricsCounterEntriesGetOK {
	return &CacheServiceMetricsCounterEntriesGetOK{}
}

/*CacheServiceMetricsCounterEntriesGetOK handles this case with default header values.

CacheServiceMetricsCounterEntriesGetOK cache service metrics counter entries get o k
*/
type CacheServiceMetricsCounterEntriesGetOK struct {
	Payload int32
}

func (o *CacheServiceMetricsCounterEntriesGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *CacheServiceMetricsCounterEntriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsCounterEntriesGetDefault creates a CacheServiceMetricsCounterEntriesGetDefault with default headers values
func NewCacheServiceMetricsCounterEntriesGetDefault(code int) *CacheServiceMetricsCounterEntriesGetDefault {
	return &CacheServiceMetricsCounterEntriesGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceMetricsCounterEntriesGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsCounterEntriesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics counter entries get default response
func (o *CacheServiceMetricsCounterEntriesGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsCounterEntriesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsCounterEntriesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsCounterEntriesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}