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

// MessagingServiceMessagesDroppedByVerGetReader is a Reader for the MessagingServiceMessagesDroppedByVerGet structure.
type MessagingServiceMessagesDroppedByVerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MessagingServiceMessagesDroppedByVerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMessagingServiceMessagesDroppedByVerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMessagingServiceMessagesDroppedByVerGetOK creates a MessagingServiceMessagesDroppedByVerGetOK with default headers values
func NewMessagingServiceMessagesDroppedByVerGetOK() *MessagingServiceMessagesDroppedByVerGetOK {
	return &MessagingServiceMessagesDroppedByVerGetOK{}
}

/*MessagingServiceMessagesDroppedByVerGetOK handles this case with default header values.

MessagingServiceMessagesDroppedByVerGetOK messaging service messages dropped by ver get o k
*/
type MessagingServiceMessagesDroppedByVerGetOK struct {
	Payload []*models.VerbCounter
}

func (o *MessagingServiceMessagesDroppedByVerGetOK) Error() string {
	return fmt.Sprintf("[GET /messaging_service/messages/dropped_by_ver][%d] messagingServiceMessagesDroppedByVerGetOK  %+v", 200, o.Payload)
}

func (o *MessagingServiceMessagesDroppedByVerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
