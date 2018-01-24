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

// GetWaitingOnFreeMemtableSpaceReader is a Reader for the GetWaitingOnFreeMemtableSpace structure.
type GetWaitingOnFreeMemtableSpaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWaitingOnFreeMemtableSpaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWaitingOnFreeMemtableSpaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWaitingOnFreeMemtableSpaceOK creates a GetWaitingOnFreeMemtableSpaceOK with default headers values
func NewGetWaitingOnFreeMemtableSpaceOK() *GetWaitingOnFreeMemtableSpaceOK {
	return &GetWaitingOnFreeMemtableSpaceOK{}
}

/*GetWaitingOnFreeMemtableSpaceOK handles this case with default header values.

GetWaitingOnFreeMemtableSpaceOK get waiting on free memtable space o k
*/
type GetWaitingOnFreeMemtableSpaceOK struct {
	Payload *models.Histogram
}

func (o *GetWaitingOnFreeMemtableSpaceOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/waiting_on_free_memtable][%d] getWaitingOnFreeMemtableSpaceOK  %+v", 200, o.Payload)
}

func (o *GetWaitingOnFreeMemtableSpaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Histogram)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
