// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla_v2/models"
)

// FindConfigRingDelayMsReader is a Reader for the FindConfigRingDelayMs structure.
type FindConfigRingDelayMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigRingDelayMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigRingDelayMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigRingDelayMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigRingDelayMsOK creates a FindConfigRingDelayMsOK with default headers values
func NewFindConfigRingDelayMsOK() *FindConfigRingDelayMsOK {
	return &FindConfigRingDelayMsOK{}
}

/*FindConfigRingDelayMsOK handles this case with default header values.

Config value
*/
type FindConfigRingDelayMsOK struct {
	Payload int64
}

func (o *FindConfigRingDelayMsOK) Error() string {
	return fmt.Sprintf("[GET /config/ring_delay_ms][%d] findConfigRingDelayMsOK  %+v", 200, o.Payload)
}

func (o *FindConfigRingDelayMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigRingDelayMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigRingDelayMsDefault creates a FindConfigRingDelayMsDefault with default headers values
func NewFindConfigRingDelayMsDefault(code int) *FindConfigRingDelayMsDefault {
	return &FindConfigRingDelayMsDefault{
		_statusCode: code,
	}
}

/*FindConfigRingDelayMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigRingDelayMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config ring delay ms default response
func (o *FindConfigRingDelayMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigRingDelayMsDefault) Error() string {
	return fmt.Sprintf("[GET /config/ring_delay_ms][%d] find_config_ring_delay_ms default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigRingDelayMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigRingDelayMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}