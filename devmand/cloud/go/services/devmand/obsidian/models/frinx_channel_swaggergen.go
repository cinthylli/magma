// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FrinxChannel A frinx channel definition
// swagger:model frinx_channel
type FrinxChannel struct {

	// authorization
	// Min Length: 1
	Authorization string `json:"authorization,omitempty"`

	// device type
	// Min Length: 1
	DeviceType string `json:"device_type,omitempty"`

	// device version
	// Min Length: 1
	DeviceVersion string `json:"device_version,omitempty"`

	// frinx port
	// Min Length: 1
	FrinxPort int32 `json:"frinx_port,omitempty"`

	// host
	// Min Length: 1
	Host string `json:"host,omitempty"`

	// password
	// Min Length: 1
	Password string `json:"password,omitempty"`

	// port
	// Min Length: 1
	Port int32 `json:"port,omitempty"`

	// transport type
	// Min Length: 1
	TransportType string `json:"transport_type,omitempty"`

	// username
	// Min Length: 1
	Username string `json:"username,omitempty"`
}

// Validate validates this frinx channel
func (m *FrinxChannel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrinxPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransportType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrinxChannel) validateAuthorization(formats strfmt.Registry) error {

	if swag.IsZero(m.Authorization) { // not required
		return nil
	}

	if err := validate.MinLength("authorization", "body", string(m.Authorization), 1); err != nil {
		return err
	}

	return nil
}

func (m *FrinxChannel) validateDeviceType(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceType) { // not required
		return nil
	}

	if err := validate.MinLength("device_type", "body", string(m.DeviceType), 1); err != nil {
		return err
	}

	return nil
}

func (m *FrinxChannel) validateDeviceVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceVersion) { // not required
		return nil
	}

	if err := validate.MinLength("device_version", "body", string(m.DeviceVersion), 1); err != nil {
		return err
	}

	return nil
}

func (m *FrinxChannel) validateFrinxPort(formats strfmt.Registry) error {

	if swag.IsZero(m.FrinxPort) { // not required
		return nil
	}

	if err := validate.MinLength("frinx_port", "body", string(m.FrinxPort), 1); err != nil {
		return err
	}

	return nil
}

func (m *FrinxChannel) validateHost(formats strfmt.Registry) error {

	if swag.IsZero(m.Host) { // not required
		return nil
	}

	if err := validate.MinLength("host", "body", string(m.Host), 1); err != nil {
		return err
	}

	return nil
}

func (m *FrinxChannel) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("password", "body", string(m.Password), 1); err != nil {
		return err
	}

	return nil
}

func (m *FrinxChannel) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinLength("port", "body", string(m.Port), 1); err != nil {
		return err
	}

	return nil
}

func (m *FrinxChannel) validateTransportType(formats strfmt.Registry) error {

	if swag.IsZero(m.TransportType) { // not required
		return nil
	}

	if err := validate.MinLength("transport_type", "body", string(m.TransportType), 1); err != nil {
		return err
	}

	return nil
}

func (m *FrinxChannel) validateUsername(formats strfmt.Registry) error {

	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if err := validate.MinLength("username", "body", string(m.Username), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrinxChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrinxChannel) UnmarshalBinary(b []byte) error {
	var res FrinxChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
