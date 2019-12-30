// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ExecutionStatus Execution Status
//
// Execution stats
// swagger:model execution-status
type ExecutionStatus string

const (

	// ExecutionStatusUnknown captures enum value "unknown"
	ExecutionStatusUnknown ExecutionStatus = "unknown"

	// ExecutionStatusQueued captures enum value "queued"
	ExecutionStatusQueued ExecutionStatus = "queued"

	// ExecutionStatusRunning captures enum value "running"
	ExecutionStatusRunning ExecutionStatus = "running"

	// ExecutionStatusCompleted captures enum value "completed"
	ExecutionStatusCompleted ExecutionStatus = "completed"

	// ExecutionStatusCancelled captures enum value "cancelled"
	ExecutionStatusCancelled ExecutionStatus = "cancelled"

	// ExecutionStatusErrored captures enum value "errored"
	ExecutionStatusErrored ExecutionStatus = "errored"
)

// for schema
var executionStatusEnum []interface{}

func init() {
	var res []ExecutionStatus
	if err := json.Unmarshal([]byte(`["unknown","queued","running","completed","cancelled","errored"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		executionStatusEnum = append(executionStatusEnum, v)
	}
}

func (m ExecutionStatus) validateExecutionStatusEnum(path, location string, value ExecutionStatus) error {
	if err := validate.Enum(path, location, value, executionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this execution status
func (m ExecutionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateExecutionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
