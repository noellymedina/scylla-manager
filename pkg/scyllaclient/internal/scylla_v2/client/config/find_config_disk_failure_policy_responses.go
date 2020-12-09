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

// FindConfigDiskFailurePolicyReader is a Reader for the FindConfigDiskFailurePolicy structure.
type FindConfigDiskFailurePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigDiskFailurePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigDiskFailurePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigDiskFailurePolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigDiskFailurePolicyOK creates a FindConfigDiskFailurePolicyOK with default headers values
func NewFindConfigDiskFailurePolicyOK() *FindConfigDiskFailurePolicyOK {
	return &FindConfigDiskFailurePolicyOK{}
}

/*FindConfigDiskFailurePolicyOK handles this case with default header values.

Config value
*/
type FindConfigDiskFailurePolicyOK struct {
	Payload string
}

func (o *FindConfigDiskFailurePolicyOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigDiskFailurePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigDiskFailurePolicyDefault creates a FindConfigDiskFailurePolicyDefault with default headers values
func NewFindConfigDiskFailurePolicyDefault(code int) *FindConfigDiskFailurePolicyDefault {
	return &FindConfigDiskFailurePolicyDefault{
		_statusCode: code,
	}
}

/*FindConfigDiskFailurePolicyDefault handles this case with default header values.

unexpected error
*/
type FindConfigDiskFailurePolicyDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config disk failure policy default response
func (o *FindConfigDiskFailurePolicyDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigDiskFailurePolicyDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigDiskFailurePolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigDiskFailurePolicyDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}