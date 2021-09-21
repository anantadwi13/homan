// Package homand provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package homand

import (
	"github.com/labstack/echo/v4"
)

// CheckHealthRes defines model for check-health-res.
type CheckHealthRes struct {
	IsAvailable bool `json:"is_available"`
}

// GeneralRes defines model for general-res.
type GeneralRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// BadRequest defines model for bad-request.
type BadRequest GeneralRes

// DefaultError defines model for default-error.
type DefaultError GeneralRes

// CheckHealthJSONBody defines parameters for CheckHealth.
type CheckHealthJSONBody struct {
	Address   string                       `json:"address"`
	CheckType CheckHealthJSONBodyCheckType `json:"check_type"`
}

// CheckHealthJSONBodyCheckType defines parameters for CheckHealth.
type CheckHealthJSONBodyCheckType string

// CheckHealthJSONRequestBody defines body for CheckHealth for application/json ContentType.
type CheckHealthJSONRequestBody CheckHealthJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /check/health)
	CheckHealth(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CheckHealth converts echo context to params.
func (w *ServerInterfaceWrapper) CheckHealth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CheckHealth(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/check/health", wrapper.CheckHealth)

}
