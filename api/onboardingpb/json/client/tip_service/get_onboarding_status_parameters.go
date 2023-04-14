// Code generated by go-swagger; DO NOT EDIT.

package tip_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetOnboardingStatusParams creates a new GetOnboardingStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOnboardingStatusParams() *GetOnboardingStatusParams {
	return &GetOnboardingStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOnboardingStatusParamsWithTimeout creates a new GetOnboardingStatusParams object
// with the ability to set a timeout on a request.
func NewGetOnboardingStatusParamsWithTimeout(timeout time.Duration) *GetOnboardingStatusParams {
	return &GetOnboardingStatusParams{
		timeout: timeout,
	}
}

// NewGetOnboardingStatusParamsWithContext creates a new GetOnboardingStatusParams object
// with the ability to set a context for a request.
func NewGetOnboardingStatusParamsWithContext(ctx context.Context) *GetOnboardingStatusParams {
	return &GetOnboardingStatusParams{
		Context: ctx,
	}
}

// NewGetOnboardingStatusParamsWithHTTPClient creates a new GetOnboardingStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOnboardingStatusParamsWithHTTPClient(client *http.Client) *GetOnboardingStatusParams {
	return &GetOnboardingStatusParams{
		HTTPClient: client,
	}
}

/*
GetOnboardingStatusParams contains all the parameters to send to the API endpoint

	for the get onboarding status operation.

	Typically these are written to a http.Request.
*/
type GetOnboardingStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get onboarding status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardingStatusParams) WithDefaults() *GetOnboardingStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get onboarding status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardingStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get onboarding status params
func (o *GetOnboardingStatusParams) WithTimeout(timeout time.Duration) *GetOnboardingStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get onboarding status params
func (o *GetOnboardingStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get onboarding status params
func (o *GetOnboardingStatusParams) WithContext(ctx context.Context) *GetOnboardingStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get onboarding status params
func (o *GetOnboardingStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get onboarding status params
func (o *GetOnboardingStatusParams) WithHTTPClient(client *http.Client) *GetOnboardingStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get onboarding status params
func (o *GetOnboardingStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOnboardingStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}