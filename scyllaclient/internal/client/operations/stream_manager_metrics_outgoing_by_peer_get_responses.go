// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StreamManagerMetricsOutgoingByPeerGetReader is a Reader for the StreamManagerMetricsOutgoingByPeerGet structure.
type StreamManagerMetricsOutgoingByPeerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamManagerMetricsOutgoingByPeerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStreamManagerMetricsOutgoingByPeerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStreamManagerMetricsOutgoingByPeerGetOK creates a StreamManagerMetricsOutgoingByPeerGetOK with default headers values
func NewStreamManagerMetricsOutgoingByPeerGetOK() *StreamManagerMetricsOutgoingByPeerGetOK {
	return &StreamManagerMetricsOutgoingByPeerGetOK{}
}

/*StreamManagerMetricsOutgoingByPeerGetOK handles this case with default header values.

StreamManagerMetricsOutgoingByPeerGetOK stream manager metrics outgoing by peer get o k
*/
type StreamManagerMetricsOutgoingByPeerGetOK struct {
	Payload int32
}

func (o *StreamManagerMetricsOutgoingByPeerGetOK) Error() string {
	return fmt.Sprintf("[GET /stream_manager/metrics/outgoing/{peer}][%d] streamManagerMetricsOutgoingByPeerGetOK  %+v", 200, o.Payload)
}

func (o *StreamManagerMetricsOutgoingByPeerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
