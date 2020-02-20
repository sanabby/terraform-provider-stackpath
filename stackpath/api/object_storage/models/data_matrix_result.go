// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataMatrixResult Time series containing a range of data points over time for each time series
// swagger:model DataMatrixResult
type DataMatrixResult struct {

	// The data points' labels
	Metric map[string]string `json:"metric,omitempty"`

	// Time series data point values
	Values []*DataMatrixResultValuesItems0 `json:"values"`
}

// Validate validates this data matrix result
func (m *DataMatrixResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataMatrixResult) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataMatrixResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataMatrixResult) UnmarshalBinary(b []byte) error {
	var res DataMatrixResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataMatrixResultValuesItems0 An individual metric data point
// swagger:model DataMatrixResultValuesItems0
type DataMatrixResultValuesItems0 struct {

	// The time that a data point was recorded
	UnixTime string `json:"unixTime,omitempty"`

	// A data point's value
	Value string `json:"value,omitempty"`
}

// Validate validates this data matrix result values items0
func (m *DataMatrixResultValuesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataMatrixResultValuesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataMatrixResultValuesItems0) UnmarshalBinary(b []byte) error {
	var res DataMatrixResultValuesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}