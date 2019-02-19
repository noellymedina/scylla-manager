// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CacheServiceRowCacheKeysToSavePostReader is a Reader for the CacheServiceRowCacheKeysToSavePost structure.
type CacheServiceRowCacheKeysToSavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceRowCacheKeysToSavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCacheServiceRowCacheKeysToSavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCacheServiceRowCacheKeysToSavePostOK creates a CacheServiceRowCacheKeysToSavePostOK with default headers values
func NewCacheServiceRowCacheKeysToSavePostOK() *CacheServiceRowCacheKeysToSavePostOK {
	return &CacheServiceRowCacheKeysToSavePostOK{}
}

/*CacheServiceRowCacheKeysToSavePostOK handles this case with default header values.

CacheServiceRowCacheKeysToSavePostOK cache service row cache keys to save post o k
*/
type CacheServiceRowCacheKeysToSavePostOK struct {
}

func (o *CacheServiceRowCacheKeysToSavePostOK) Error() string {
	return fmt.Sprintf("[POST /cache_service/row_cache_keys_to_save][%d] cacheServiceRowCacheKeysToSavePostOK ", 200)
}

func (o *CacheServiceRowCacheKeysToSavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
