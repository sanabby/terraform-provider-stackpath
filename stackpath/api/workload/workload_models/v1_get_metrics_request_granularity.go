// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1GetMetricsRequestGranularity The time granularity of retrieved metrics
//
// This field does not apply when retrieving INSTANCE type metrics
//
//  - DEFAULT: The current default is PT5M
//  - PT5M: Metrics every five minutes
//  - PT1H: Metrics every hour
//  - P1D: Metrics every day
//  - PT1M: Metrics every one minute
// swagger:model v1GetMetricsRequestGranularity
type V1GetMetricsRequestGranularity string

const (

	// V1GetMetricsRequestGranularityDEFAULT captures enum value "DEFAULT"
	V1GetMetricsRequestGranularityDEFAULT V1GetMetricsRequestGranularity = "DEFAULT"

	// V1GetMetricsRequestGranularityPT5M captures enum value "PT5M"
	V1GetMetricsRequestGranularityPT5M V1GetMetricsRequestGranularity = "PT5M"

	// V1GetMetricsRequestGranularityPT1H captures enum value "PT1H"
	V1GetMetricsRequestGranularityPT1H V1GetMetricsRequestGranularity = "PT1H"

	// V1GetMetricsRequestGranularityP1D captures enum value "P1D"
	V1GetMetricsRequestGranularityP1D V1GetMetricsRequestGranularity = "P1D"

	// V1GetMetricsRequestGranularityPT1M captures enum value "PT1M"
	V1GetMetricsRequestGranularityPT1M V1GetMetricsRequestGranularity = "PT1M"
)

// for schema
var v1GetMetricsRequestGranularityEnum []interface{}

func init() {
	var res []V1GetMetricsRequestGranularity
	if err := json.Unmarshal([]byte(`["DEFAULT","PT5M","PT1H","P1D","PT1M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1GetMetricsRequestGranularityEnum = append(v1GetMetricsRequestGranularityEnum, v)
	}
}

func (m V1GetMetricsRequestGranularity) validateV1GetMetricsRequestGranularityEnum(path, location string, value V1GetMetricsRequestGranularity) error {
	if err := validate.Enum(path, location, value, v1GetMetricsRequestGranularityEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 get metrics request granularity
func (m V1GetMetricsRequestGranularity) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1GetMetricsRequestGranularityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
