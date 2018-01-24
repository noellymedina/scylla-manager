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

// GetDroppedMessagesByVerReader is a Reader for the GetDroppedMessagesByVer structure.
type GetDroppedMessagesByVerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDroppedMessagesByVerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDroppedMessagesByVerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDroppedMessagesByVerOK creates a GetDroppedMessagesByVerOK with default headers values
func NewGetDroppedMessagesByVerOK() *GetDroppedMessagesByVerOK {
	return &GetDroppedMessagesByVerOK{}
}

/*GetDroppedMessagesByVerOK handles this case with default header values.

GetDroppedMessagesByVerOK get dropped messages by ver o k
*/
type GetDroppedMessagesByVerOK struct {
	Payload models.GetDroppedMessagesByVerOKBody
}

func (o *GetDroppedMessagesByVerOK) Error() string {
	return fmt.Sprintf("[GET /messaging_service/messages/dropped_by_ver][%d] getDroppedMessagesByVerOK  %+v", 200, o.Payload)
}

func (o *GetDroppedMessagesByVerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
