// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/models"
)

// GetColumnFamilyReader is a Reader for the GetColumnFamily structure.
type GetColumnFamilyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetColumnFamilyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetColumnFamilyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetColumnFamilyOK creates a GetColumnFamilyOK with default headers values
func NewGetColumnFamilyOK() *GetColumnFamilyOK {
	return &GetColumnFamilyOK{}
}

/*GetColumnFamilyOK handles this case with default header values.

GetColumnFamilyOK get column family o k
*/
type GetColumnFamilyOK struct {
	Payload models.GetColumnFamilyOKBody
}

func (o *GetColumnFamilyOK) Error() string {
	return fmt.Sprintf("[GET /column_family/][%d] getColumnFamilyOK  %+v", 200, o.Payload)
}

func (o *GetColumnFamilyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
