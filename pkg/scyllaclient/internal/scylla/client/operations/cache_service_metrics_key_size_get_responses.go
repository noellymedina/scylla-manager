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

	models "github.com/scylladb/mermaid/pkg/scyllaclient/internal/scylla/models"
)

// CacheServiceMetricsKeySizeGetReader is a Reader for the CacheServiceMetricsKeySizeGet structure.
type CacheServiceMetricsKeySizeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsKeySizeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsKeySizeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsKeySizeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsKeySizeGetOK creates a CacheServiceMetricsKeySizeGetOK with default headers values
func NewCacheServiceMetricsKeySizeGetOK() *CacheServiceMetricsKeySizeGetOK {
	return &CacheServiceMetricsKeySizeGetOK{}
}

/*CacheServiceMetricsKeySizeGetOK handles this case with default header values.

CacheServiceMetricsKeySizeGetOK cache service metrics key size get o k
*/
type CacheServiceMetricsKeySizeGetOK struct {
	Payload interface{}
}

func (o *CacheServiceMetricsKeySizeGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CacheServiceMetricsKeySizeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsKeySizeGetDefault creates a CacheServiceMetricsKeySizeGetDefault with default headers values
func NewCacheServiceMetricsKeySizeGetDefault(code int) *CacheServiceMetricsKeySizeGetDefault {
	return &CacheServiceMetricsKeySizeGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceMetricsKeySizeGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsKeySizeGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics key size get default response
func (o *CacheServiceMetricsKeySizeGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsKeySizeGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsKeySizeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsKeySizeGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}