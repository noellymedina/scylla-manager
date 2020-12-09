// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla/models"
)

// SystemLoggerPostReader is a Reader for the SystemLoggerPost structure.
type SystemLoggerPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemLoggerPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemLoggerPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSystemLoggerPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSystemLoggerPostOK creates a SystemLoggerPostOK with default headers values
func NewSystemLoggerPostOK() *SystemLoggerPostOK {
	return &SystemLoggerPostOK{}
}

/*SystemLoggerPostOK handles this case with default header values.

SystemLoggerPostOK system logger post o k
*/
type SystemLoggerPostOK struct {
}

func (o *SystemLoggerPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSystemLoggerPostDefault creates a SystemLoggerPostDefault with default headers values
func NewSystemLoggerPostDefault(code int) *SystemLoggerPostDefault {
	return &SystemLoggerPostDefault{
		_statusCode: code,
	}
}

/*SystemLoggerPostDefault handles this case with default header values.

internal server error
*/
type SystemLoggerPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the system logger post default response
func (o *SystemLoggerPostDefault) Code() int {
	return o._statusCode
}

func (o *SystemLoggerPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *SystemLoggerPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *SystemLoggerPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}