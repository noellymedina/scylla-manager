// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v1/models"
)

// MessagingServiceMessagesDroppedGetReader is a Reader for the MessagingServiceMessagesDroppedGet structure.
type MessagingServiceMessagesDroppedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MessagingServiceMessagesDroppedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMessagingServiceMessagesDroppedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMessagingServiceMessagesDroppedGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMessagingServiceMessagesDroppedGetOK creates a MessagingServiceMessagesDroppedGetOK with default headers values
func NewMessagingServiceMessagesDroppedGetOK() *MessagingServiceMessagesDroppedGetOK {
	return &MessagingServiceMessagesDroppedGetOK{}
}

/*MessagingServiceMessagesDroppedGetOK handles this case with default header values.

MessagingServiceMessagesDroppedGetOK messaging service messages dropped get o k
*/
type MessagingServiceMessagesDroppedGetOK struct {
	Payload []*models.MessageCounter
}

func (o *MessagingServiceMessagesDroppedGetOK) GetPayload() []*models.MessageCounter {
	return o.Payload
}

func (o *MessagingServiceMessagesDroppedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMessagingServiceMessagesDroppedGetDefault creates a MessagingServiceMessagesDroppedGetDefault with default headers values
func NewMessagingServiceMessagesDroppedGetDefault(code int) *MessagingServiceMessagesDroppedGetDefault {
	return &MessagingServiceMessagesDroppedGetDefault{
		_statusCode: code,
	}
}

/*MessagingServiceMessagesDroppedGetDefault handles this case with default header values.

internal server error
*/
type MessagingServiceMessagesDroppedGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the messaging service messages dropped get default response
func (o *MessagingServiceMessagesDroppedGetDefault) Code() int {
	return o._statusCode
}

func (o *MessagingServiceMessagesDroppedGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *MessagingServiceMessagesDroppedGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *MessagingServiceMessagesDroppedGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}