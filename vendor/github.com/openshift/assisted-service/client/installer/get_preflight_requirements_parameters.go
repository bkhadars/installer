// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewGetPreflightRequirementsParams creates a new GetPreflightRequirementsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPreflightRequirementsParams() *GetPreflightRequirementsParams {
	return &GetPreflightRequirementsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPreflightRequirementsParamsWithTimeout creates a new GetPreflightRequirementsParams object
// with the ability to set a timeout on a request.
func NewGetPreflightRequirementsParamsWithTimeout(timeout time.Duration) *GetPreflightRequirementsParams {
	return &GetPreflightRequirementsParams{
		timeout: timeout,
	}
}

// NewGetPreflightRequirementsParamsWithContext creates a new GetPreflightRequirementsParams object
// with the ability to set a context for a request.
func NewGetPreflightRequirementsParamsWithContext(ctx context.Context) *GetPreflightRequirementsParams {
	return &GetPreflightRequirementsParams{
		Context: ctx,
	}
}

// NewGetPreflightRequirementsParamsWithHTTPClient creates a new GetPreflightRequirementsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPreflightRequirementsParamsWithHTTPClient(client *http.Client) *GetPreflightRequirementsParams {
	return &GetPreflightRequirementsParams{
		HTTPClient: client,
	}
}

/* GetPreflightRequirementsParams contains all the parameters to send to the API endpoint
   for the get preflight requirements operation.

   Typically these are written to a http.Request.
*/
type GetPreflightRequirementsParams struct {

	/* ClusterID.

	   The cluster to return preflight requrements for.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get preflight requirements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPreflightRequirementsParams) WithDefaults() *GetPreflightRequirementsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get preflight requirements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPreflightRequirementsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get preflight requirements params
func (o *GetPreflightRequirementsParams) WithTimeout(timeout time.Duration) *GetPreflightRequirementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get preflight requirements params
func (o *GetPreflightRequirementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get preflight requirements params
func (o *GetPreflightRequirementsParams) WithContext(ctx context.Context) *GetPreflightRequirementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get preflight requirements params
func (o *GetPreflightRequirementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get preflight requirements params
func (o *GetPreflightRequirementsParams) WithHTTPClient(client *http.Client) *GetPreflightRequirementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get preflight requirements params
func (o *GetPreflightRequirementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get preflight requirements params
func (o *GetPreflightRequirementsParams) WithClusterID(clusterID strfmt.UUID) *GetPreflightRequirementsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get preflight requirements params
func (o *GetPreflightRequirementsParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPreflightRequirementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
