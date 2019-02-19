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

// ColumnFamilyMinimumCompactionByNamePostReader is a Reader for the ColumnFamilyMinimumCompactionByNamePost structure.
type ColumnFamilyMinimumCompactionByNamePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMinimumCompactionByNamePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMinimumCompactionByNamePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMinimumCompactionByNamePostOK creates a ColumnFamilyMinimumCompactionByNamePostOK with default headers values
func NewColumnFamilyMinimumCompactionByNamePostOK() *ColumnFamilyMinimumCompactionByNamePostOK {
	return &ColumnFamilyMinimumCompactionByNamePostOK{}
}

/*ColumnFamilyMinimumCompactionByNamePostOK handles this case with default header values.

ColumnFamilyMinimumCompactionByNamePostOK column family minimum compaction by name post o k
*/
type ColumnFamilyMinimumCompactionByNamePostOK struct {
	Payload string
}

func (o *ColumnFamilyMinimumCompactionByNamePostOK) Error() string {
	return fmt.Sprintf("[POST /column_family/minimum_compaction/{name}][%d] columnFamilyMinimumCompactionByNamePostOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMinimumCompactionByNamePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
