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

// ColumnFamilyMetricsMemtableSwitchCountByNameGetReader is a Reader for the ColumnFamilyMetricsMemtableSwitchCountByNameGet structure.
type ColumnFamilyMetricsMemtableSwitchCountByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMemtableSwitchCountByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMemtableSwitchCountByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsMemtableSwitchCountByNameGetOK creates a ColumnFamilyMetricsMemtableSwitchCountByNameGetOK with default headers values
func NewColumnFamilyMetricsMemtableSwitchCountByNameGetOK() *ColumnFamilyMetricsMemtableSwitchCountByNameGetOK {
	return &ColumnFamilyMetricsMemtableSwitchCountByNameGetOK{}
}

/*ColumnFamilyMetricsMemtableSwitchCountByNameGetOK handles this case with default header values.

ColumnFamilyMetricsMemtableSwitchCountByNameGetOK column family metrics memtable switch count by name get o k
*/
type ColumnFamilyMetricsMemtableSwitchCountByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsMemtableSwitchCountByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/memtable_switch_count/{name}][%d] columnFamilyMetricsMemtableSwitchCountByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsMemtableSwitchCountByNameGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableSwitchCountByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}