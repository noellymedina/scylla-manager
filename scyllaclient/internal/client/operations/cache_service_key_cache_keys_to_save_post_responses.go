// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CacheServiceKeyCacheKeysToSavePostReader is a Reader for the CacheServiceKeyCacheKeysToSavePost structure.
type CacheServiceKeyCacheKeysToSavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceKeyCacheKeysToSavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCacheServiceKeyCacheKeysToSavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCacheServiceKeyCacheKeysToSavePostOK creates a CacheServiceKeyCacheKeysToSavePostOK with default headers values
func NewCacheServiceKeyCacheKeysToSavePostOK() *CacheServiceKeyCacheKeysToSavePostOK {
	return &CacheServiceKeyCacheKeysToSavePostOK{}
}

/*CacheServiceKeyCacheKeysToSavePostOK handles this case with default header values.

CacheServiceKeyCacheKeysToSavePostOK cache service key cache keys to save post o k
*/
type CacheServiceKeyCacheKeysToSavePostOK struct {
}

func (o *CacheServiceKeyCacheKeysToSavePostOK) Error() string {
	return fmt.Sprintf("[POST /cache_service/key_cache_keys_to_save][%d] cacheServiceKeyCacheKeysToSavePostOK ", 200)
}

func (o *CacheServiceKeyCacheKeysToSavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
