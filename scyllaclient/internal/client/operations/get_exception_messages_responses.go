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

// GetExceptionMessagesReader is a Reader for the GetExceptionMessages structure.
type GetExceptionMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExceptionMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetExceptionMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetExceptionMessagesOK creates a GetExceptionMessagesOK with default headers values
func NewGetExceptionMessagesOK() *GetExceptionMessagesOK {
	return &GetExceptionMessagesOK{}
}

/*GetExceptionMessagesOK handles this case with default header values.

GetExceptionMessagesOK get exception messages o k
*/
type GetExceptionMessagesOK struct {
	Payload models.GetExceptionMessagesOKBody
}

func (o *GetExceptionMessagesOK) Error() string {
	return fmt.Sprintf("[GET /messaging_service/messages/exception][%d] getExceptionMessagesOK  %+v", 200, o.Payload)
}

func (o *GetExceptionMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
