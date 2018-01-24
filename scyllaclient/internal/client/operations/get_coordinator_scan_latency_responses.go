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

// GetCoordinatorScanLatencyReader is a Reader for the GetCoordinatorScanLatency structure.
type GetCoordinatorScanLatencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCoordinatorScanLatencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCoordinatorScanLatencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCoordinatorScanLatencyOK creates a GetCoordinatorScanLatencyOK with default headers values
func NewGetCoordinatorScanLatencyOK() *GetCoordinatorScanLatencyOK {
	return &GetCoordinatorScanLatencyOK{}
}

/*GetCoordinatorScanLatencyOK handles this case with default header values.

GetCoordinatorScanLatencyOK get coordinator scan latency o k
*/
type GetCoordinatorScanLatencyOK struct {
	Payload *models.Histogram
}

func (o *GetCoordinatorScanLatencyOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/coordinator/scan][%d] getCoordinatorScanLatencyOK  %+v", 200, o.Payload)
}

func (o *GetCoordinatorScanLatencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Histogram)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
