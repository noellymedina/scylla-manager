// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CacheServiceCounterCacheSavePeriodGetReader is a Reader for the CacheServiceCounterCacheSavePeriodGet structure.
type CacheServiceCounterCacheSavePeriodGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceCounterCacheSavePeriodGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCacheServiceCounterCacheSavePeriodGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCacheServiceCounterCacheSavePeriodGetOK creates a CacheServiceCounterCacheSavePeriodGetOK with default headers values
func NewCacheServiceCounterCacheSavePeriodGetOK() *CacheServiceCounterCacheSavePeriodGetOK {
	return &CacheServiceCounterCacheSavePeriodGetOK{}
}

/*CacheServiceCounterCacheSavePeriodGetOK handles this case with default header values.

CacheServiceCounterCacheSavePeriodGetOK cache service counter cache save period get o k
*/
type CacheServiceCounterCacheSavePeriodGetOK struct {
	Payload int32
}

func (o *CacheServiceCounterCacheSavePeriodGetOK) Error() string {
	return fmt.Sprintf("[GET /cache_service/counter_cache_save_period][%d] cacheServiceCounterCacheSavePeriodGetOK  %+v", 200, o.Payload)
}

func (o *CacheServiceCounterCacheSavePeriodGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
