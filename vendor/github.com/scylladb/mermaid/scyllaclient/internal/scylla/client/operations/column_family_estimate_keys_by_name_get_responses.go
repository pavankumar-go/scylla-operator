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

// ColumnFamilyEstimateKeysByNameGetReader is a Reader for the ColumnFamilyEstimateKeysByNameGet structure.
type ColumnFamilyEstimateKeysByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyEstimateKeysByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyEstimateKeysByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyEstimateKeysByNameGetOK creates a ColumnFamilyEstimateKeysByNameGetOK with default headers values
func NewColumnFamilyEstimateKeysByNameGetOK() *ColumnFamilyEstimateKeysByNameGetOK {
	return &ColumnFamilyEstimateKeysByNameGetOK{}
}

/*ColumnFamilyEstimateKeysByNameGetOK handles this case with default header values.

ColumnFamilyEstimateKeysByNameGetOK column family estimate keys by name get o k
*/
type ColumnFamilyEstimateKeysByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyEstimateKeysByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/estimate_keys/{name}][%d] columnFamilyEstimateKeysByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyEstimateKeysByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyEstimateKeysByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}