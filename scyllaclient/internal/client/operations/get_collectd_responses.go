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

// GetCollectdReader is a Reader for the GetCollectd structure.
type GetCollectdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCollectdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCollectdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCollectdOK creates a GetCollectdOK with default headers values
func NewGetCollectdOK() *GetCollectdOK {
	return &GetCollectdOK{}
}

/*GetCollectdOK handles this case with default header values.

GetCollectdOK get collectd o k
*/
type GetCollectdOK struct {
	Payload models.GetCollectdOKBody
}

func (o *GetCollectdOK) Error() string {
	return fmt.Sprintf("[GET /collectd/{pluginid}][%d] getCollectdOK  %+v", 200, o.Payload)
}

func (o *GetCollectdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
