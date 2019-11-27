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

// ColumnFamilyDroppableRatioByNameGetReader is a Reader for the ColumnFamilyDroppableRatioByNameGet structure.
type ColumnFamilyDroppableRatioByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyDroppableRatioByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyDroppableRatioByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyDroppableRatioByNameGetOK creates a ColumnFamilyDroppableRatioByNameGetOK with default headers values
func NewColumnFamilyDroppableRatioByNameGetOK() *ColumnFamilyDroppableRatioByNameGetOK {
	return &ColumnFamilyDroppableRatioByNameGetOK{}
}

/*ColumnFamilyDroppableRatioByNameGetOK handles this case with default header values.

ColumnFamilyDroppableRatioByNameGetOK column family droppable ratio by name get o k
*/
type ColumnFamilyDroppableRatioByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyDroppableRatioByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/droppable_ratio/{name}][%d] columnFamilyDroppableRatioByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyDroppableRatioByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyDroppableRatioByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}