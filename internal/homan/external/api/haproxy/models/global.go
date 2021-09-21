// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Global Global
//
// HAProxy global configuration
//
// swagger:model global
type Global struct {

	// CPU maps
	CPUMaps []*CPUMap `json:"cpu_maps"`

	// runtime a p is
	RuntimeAPIs []*RuntimeAPI `json:"runtime_apis"`

	// chroot
	// Pattern: ^[^\s]+$
	Chroot string `json:"chroot,omitempty"`

	// daemon
	// Enum: [enabled disabled]
	Daemon string `json:"daemon,omitempty"`

	// external check
	ExternalCheck bool `json:"external_check,omitempty"`

	// group
	// Pattern: ^[^\s]+$
	Group string `json:"group,omitempty"`

	// hard stop after
	HardStopAfter *int64 `json:"hard_stop_after,omitempty"`

	// localpeer
	// Pattern: ^[^\s]+$
	Localpeer string `json:"localpeer,omitempty"`

	// log send hostname
	LogSendHostname *GlobalLogSendHostname `json:"log_send_hostname,omitempty"`

	// lua loads
	LuaLoads []*LuaLoad `json:"lua_loads"`

	// lua prepend path
	LuaPrependPath []*LuaPrependPath `json:"lua_prepend_path"`

	// master worker
	MasterWorker bool `json:"master-worker,omitempty"`

	// maxconn
	Maxconn int64 `json:"maxconn,omitempty"`

	// nbproc
	Nbproc int64 `json:"nbproc,omitempty"`

	// nbthread
	Nbthread int64 `json:"nbthread,omitempty"`

	// pidfile
	Pidfile string `json:"pidfile,omitempty"`

	// server state base
	// Pattern: ^[^\s]+$
	ServerStateBase string `json:"server_state_base,omitempty"`

	// server state file
	// Pattern: ^[^\s]+$
	ServerStateFile string `json:"server_state_file,omitempty"`

	// ssl default bind ciphers
	SslDefaultBindCiphers string `json:"ssl_default_bind_ciphers,omitempty"`

	// ssl default bind ciphersuites
	SslDefaultBindCiphersuites string `json:"ssl_default_bind_ciphersuites,omitempty"`

	// ssl default bind options
	SslDefaultBindOptions string `json:"ssl_default_bind_options,omitempty"`

	// ssl default server ciphers
	SslDefaultServerCiphers string `json:"ssl_default_server_ciphers,omitempty"`

	// ssl default server ciphersuites
	SslDefaultServerCiphersuites string `json:"ssl_default_server_ciphersuites,omitempty"`

	// ssl default server options
	SslDefaultServerOptions string `json:"ssl_default_server_options,omitempty"`

	// ssl mode async
	// Enum: [enabled disabled]
	SslModeAsync string `json:"ssl_mode_async,omitempty"`

	// stats timeout
	StatsTimeout *int64 `json:"stats_timeout,omitempty"`

	// tune ssl default dh param
	TuneSslDefaultDhParam int64 `json:"tune_ssl_default_dh_param,omitempty"`

	// user
	// Pattern: ^[^\s]+$
	User string `json:"user,omitempty"`
}

// Validate validates this global
func (m *Global) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUMaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntimeAPIs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChroot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaemon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalpeer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogSendHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuaLoads(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuaPrependPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerStateBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerStateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslModeAsync(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Global) validateCPUMaps(formats strfmt.Registry) error {
	if swag.IsZero(m.CPUMaps) { // not required
		return nil
	}

	for i := 0; i < len(m.CPUMaps); i++ {
		if swag.IsZero(m.CPUMaps[i]) { // not required
			continue
		}

		if m.CPUMaps[i] != nil {
			if err := m.CPUMaps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpu_maps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateRuntimeAPIs(formats strfmt.Registry) error {
	if swag.IsZero(m.RuntimeAPIs) { // not required
		return nil
	}

	for i := 0; i < len(m.RuntimeAPIs); i++ {
		if swag.IsZero(m.RuntimeAPIs[i]) { // not required
			continue
		}

		if m.RuntimeAPIs[i] != nil {
			if err := m.RuntimeAPIs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runtime_apis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateChroot(formats strfmt.Registry) error {
	if swag.IsZero(m.Chroot) { // not required
		return nil
	}

	if err := validate.Pattern("chroot", "body", m.Chroot, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var globalTypeDaemonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalTypeDaemonPropEnum = append(globalTypeDaemonPropEnum, v)
	}
}

const (

	// GlobalDaemonEnabled captures enum value "enabled"
	GlobalDaemonEnabled string = "enabled"

	// GlobalDaemonDisabled captures enum value "disabled"
	GlobalDaemonDisabled string = "disabled"
)

// prop value enum
func (m *Global) validateDaemonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, globalTypeDaemonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Global) validateDaemon(formats strfmt.Registry) error {
	if swag.IsZero(m.Daemon) { // not required
		return nil
	}

	// value enum
	if err := m.validateDaemonEnum("daemon", "body", m.Daemon); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if err := validate.Pattern("group", "body", m.Group, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateLocalpeer(formats strfmt.Registry) error {
	if swag.IsZero(m.Localpeer) { // not required
		return nil
	}

	if err := validate.Pattern("localpeer", "body", m.Localpeer, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateLogSendHostname(formats strfmt.Registry) error {
	if swag.IsZero(m.LogSendHostname) { // not required
		return nil
	}

	if m.LogSendHostname != nil {
		if err := m.LogSendHostname.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_send_hostname")
			}
			return err
		}
	}

	return nil
}

func (m *Global) validateLuaLoads(formats strfmt.Registry) error {
	if swag.IsZero(m.LuaLoads) { // not required
		return nil
	}

	for i := 0; i < len(m.LuaLoads); i++ {
		if swag.IsZero(m.LuaLoads[i]) { // not required
			continue
		}

		if m.LuaLoads[i] != nil {
			if err := m.LuaLoads[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lua_loads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateLuaPrependPath(formats strfmt.Registry) error {
	if swag.IsZero(m.LuaPrependPath) { // not required
		return nil
	}

	for i := 0; i < len(m.LuaPrependPath); i++ {
		if swag.IsZero(m.LuaPrependPath[i]) { // not required
			continue
		}

		if m.LuaPrependPath[i] != nil {
			if err := m.LuaPrependPath[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lua_prepend_path" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) validateServerStateBase(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerStateBase) { // not required
		return nil
	}

	if err := validate.Pattern("server_state_base", "body", m.ServerStateBase, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateServerStateFile(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerStateFile) { // not required
		return nil
	}

	if err := validate.Pattern("server_state_file", "body", m.ServerStateFile, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var globalTypeSslModeAsyncPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalTypeSslModeAsyncPropEnum = append(globalTypeSslModeAsyncPropEnum, v)
	}
}

const (

	// GlobalSslModeAsyncEnabled captures enum value "enabled"
	GlobalSslModeAsyncEnabled string = "enabled"

	// GlobalSslModeAsyncDisabled captures enum value "disabled"
	GlobalSslModeAsyncDisabled string = "disabled"
)

// prop value enum
func (m *Global) validateSslModeAsyncEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, globalTypeSslModeAsyncPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Global) validateSslModeAsync(formats strfmt.Registry) error {
	if swag.IsZero(m.SslModeAsync) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslModeAsyncEnum("ssl_mode_async", "body", m.SslModeAsync); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if err := validate.Pattern("user", "body", m.User, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this global based on the context it is used
func (m *Global) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCPUMaps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuntimeAPIs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogSendHostname(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLuaLoads(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLuaPrependPath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Global) contextValidateCPUMaps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CPUMaps); i++ {

		if m.CPUMaps[i] != nil {
			if err := m.CPUMaps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpu_maps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) contextValidateRuntimeAPIs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RuntimeAPIs); i++ {

		if m.RuntimeAPIs[i] != nil {
			if err := m.RuntimeAPIs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runtime_apis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) contextValidateLogSendHostname(ctx context.Context, formats strfmt.Registry) error {

	if m.LogSendHostname != nil {
		if err := m.LogSendHostname.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_send_hostname")
			}
			return err
		}
	}

	return nil
}

func (m *Global) contextValidateLuaLoads(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LuaLoads); i++ {

		if m.LuaLoads[i] != nil {
			if err := m.LuaLoads[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lua_loads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Global) contextValidateLuaPrependPath(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LuaPrependPath); i++ {

		if m.LuaPrependPath[i] != nil {
			if err := m.LuaPrependPath[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lua_prepend_path" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Global) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Global) UnmarshalBinary(b []byte) error {
	var res Global
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CPUMap CPU map
//
// swagger:model CPUMap
type CPUMap struct {

	// cpu set
	// Required: true
	CPUSet *string `json:"cpu_set"`

	// process
	// Required: true
	Process *string `json:"process"`
}

// Validate validates this CPU map
func (m *CPUMap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CPUMap) validateCPUSet(formats strfmt.Registry) error {

	if err := validate.Required("cpu_set", "body", m.CPUSet); err != nil {
		return err
	}

	return nil
}

func (m *CPUMap) validateProcess(formats strfmt.Registry) error {

	if err := validate.Required("process", "body", m.Process); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this CPU map based on context it is used
func (m *CPUMap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CPUMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPUMap) UnmarshalBinary(b []byte) error {
	var res CPUMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GlobalLogSendHostname global log send hostname
//
// swagger:model GlobalLogSendHostname
type GlobalLogSendHostname struct {

	// enabled
	// Required: true
	// Enum: [enabled disabled]
	Enabled *string `json:"enabled"`

	// param
	// Pattern: ^[^\s]+$
	Param string `json:"param,omitempty"`
}

// Validate validates this global log send hostname
func (m *GlobalLogSendHostname) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParam(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var globalLogSendHostnameTypeEnabledPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalLogSendHostnameTypeEnabledPropEnum = append(globalLogSendHostnameTypeEnabledPropEnum, v)
	}
}

const (

	// GlobalLogSendHostnameEnabledEnabled captures enum value "enabled"
	GlobalLogSendHostnameEnabledEnabled string = "enabled"

	// GlobalLogSendHostnameEnabledDisabled captures enum value "disabled"
	GlobalLogSendHostnameEnabledDisabled string = "disabled"
)

// prop value enum
func (m *GlobalLogSendHostname) validateEnabledEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, globalLogSendHostnameTypeEnabledPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GlobalLogSendHostname) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("log_send_hostname"+"."+"enabled", "body", m.Enabled); err != nil {
		return err
	}

	// value enum
	if err := m.validateEnabledEnum("log_send_hostname"+"."+"enabled", "body", *m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *GlobalLogSendHostname) validateParam(formats strfmt.Registry) error {
	if swag.IsZero(m.Param) { // not required
		return nil
	}

	if err := validate.Pattern("log_send_hostname"+"."+"param", "body", m.Param, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this global log send hostname based on context it is used
func (m *GlobalLogSendHostname) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GlobalLogSendHostname) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalLogSendHostname) UnmarshalBinary(b []byte) error {
	var res GlobalLogSendHostname
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LuaLoad lua load
//
// swagger:model LuaLoad
type LuaLoad struct {

	// file
	// Required: true
	// Pattern: ^[^\s]+$
	File *string `json:"file"`
}

// Validate validates this lua load
func (m *LuaLoad) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LuaLoad) validateFile(formats strfmt.Registry) error {

	if err := validate.Required("file", "body", m.File); err != nil {
		return err
	}

	if err := validate.Pattern("file", "body", *m.File, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this lua load based on context it is used
func (m *LuaLoad) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LuaLoad) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LuaLoad) UnmarshalBinary(b []byte) error {
	var res LuaLoad
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LuaPrependPath lua prepend path
//
// swagger:model LuaPrependPath
type LuaPrependPath struct {

	// path
	// Required: true
	// Pattern: ^[^\s]+$
	Path *string `json:"path"`

	// type
	// Enum: [path cpath]
	Type string `json:"type,omitempty"`
}

// Validate validates this lua prepend path
func (m *LuaPrependPath) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LuaPrependPath) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	if err := validate.Pattern("path", "body", *m.Path, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var luaPrependPathTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["path","cpath"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		luaPrependPathTypeTypePropEnum = append(luaPrependPathTypeTypePropEnum, v)
	}
}

const (

	// LuaPrependPathTypePath captures enum value "path"
	LuaPrependPathTypePath string = "path"

	// LuaPrependPathTypeCpath captures enum value "cpath"
	LuaPrependPathTypeCpath string = "cpath"
)

// prop value enum
func (m *LuaPrependPath) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, luaPrependPathTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LuaPrependPath) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this lua prepend path based on context it is used
func (m *LuaPrependPath) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LuaPrependPath) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LuaPrependPath) UnmarshalBinary(b []byte) error {
	var res LuaPrependPath
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RuntimeAPI runtime API
//
// swagger:model RuntimeAPI
type RuntimeAPI struct {

	// address
	// Required: true
	// Pattern: ^[^\s]+$
	Address *string `json:"address"`

	// expose fd listeners
	ExposeFdListeners bool `json:"exposeFdListeners,omitempty"`

	// level
	// Enum: [user operator admin]
	Level string `json:"level,omitempty"`

	// mode
	// Pattern: ^[^\s]+$
	Mode string `json:"mode,omitempty"`

	// process
	// Pattern: ^[^\s]+$
	Process string `json:"process,omitempty"`
}

// Validate validates this runtime API
func (m *RuntimeAPI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuntimeAPI) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	if err := validate.Pattern("address", "body", *m.Address, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var runtimeApiTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","operator","admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		runtimeApiTypeLevelPropEnum = append(runtimeApiTypeLevelPropEnum, v)
	}
}

const (

	// RuntimeAPILevelUser captures enum value "user"
	RuntimeAPILevelUser string = "user"

	// RuntimeAPILevelOperator captures enum value "operator"
	RuntimeAPILevelOperator string = "operator"

	// RuntimeAPILevelAdmin captures enum value "admin"
	RuntimeAPILevelAdmin string = "admin"
)

// prop value enum
func (m *RuntimeAPI) validateLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, runtimeApiTypeLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RuntimeAPI) validateLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeAPI) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := validate.Pattern("mode", "body", m.Mode, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *RuntimeAPI) validateProcess(formats strfmt.Registry) error {
	if swag.IsZero(m.Process) { // not required
		return nil
	}

	if err := validate.Pattern("process", "body", m.Process, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this runtime API based on context it is used
func (m *RuntimeAPI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RuntimeAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuntimeAPI) UnmarshalBinary(b []byte) error {
	var res RuntimeAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}