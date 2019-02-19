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

// CacheServiceRowCacheSavePeriodGetReader is a Reader for the CacheServiceRowCacheSavePeriodGet structure.
type CacheServiceRowCacheSavePeriodGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceRowCacheSavePeriodGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCacheServiceRowCacheSavePeriodGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCacheServiceRowCacheSavePeriodGetOK creates a CacheServiceRowCacheSavePeriodGetOK with default headers values
func NewCacheServiceRowCacheSavePeriodGetOK() *CacheServiceRowCacheSavePeriodGetOK {
	return &CacheServiceRowCacheSavePeriodGetOK{}
}

/*CacheServiceRowCacheSavePeriodGetOK handles this case with default header values.

CacheServiceRowCacheSavePeriodGetOK cache service row cache save period get o k
*/
type CacheServiceRowCacheSavePeriodGetOK struct {
	Payload int32
}

func (o *CacheServiceRowCacheSavePeriodGetOK) Error() string {
	return fmt.Sprintf("[GET /cache_service/row_cache_save_period][%d] cacheServiceRowCacheSavePeriodGetOK  %+v", 200, o.Payload)
}

func (o *CacheServiceRowCacheSavePeriodGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
