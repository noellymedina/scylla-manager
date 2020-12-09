// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigRowCacheSavePeriodReader is a Reader for the FindConfigRowCacheSavePeriod structure.
type FindConfigRowCacheSavePeriodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigRowCacheSavePeriodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigRowCacheSavePeriodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigRowCacheSavePeriodDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigRowCacheSavePeriodOK creates a FindConfigRowCacheSavePeriodOK with default headers values
func NewFindConfigRowCacheSavePeriodOK() *FindConfigRowCacheSavePeriodOK {
	return &FindConfigRowCacheSavePeriodOK{}
}

/*FindConfigRowCacheSavePeriodOK handles this case with default header values.

Config value
*/
type FindConfigRowCacheSavePeriodOK struct {
	Payload int64
}

func (o *FindConfigRowCacheSavePeriodOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigRowCacheSavePeriodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigRowCacheSavePeriodDefault creates a FindConfigRowCacheSavePeriodDefault with default headers values
func NewFindConfigRowCacheSavePeriodDefault(code int) *FindConfigRowCacheSavePeriodDefault {
	return &FindConfigRowCacheSavePeriodDefault{
		_statusCode: code,
	}
}

/*FindConfigRowCacheSavePeriodDefault handles this case with default header values.

unexpected error
*/
type FindConfigRowCacheSavePeriodDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config row cache save period default response
func (o *FindConfigRowCacheSavePeriodDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigRowCacheSavePeriodDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigRowCacheSavePeriodDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigRowCacheSavePeriodDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}