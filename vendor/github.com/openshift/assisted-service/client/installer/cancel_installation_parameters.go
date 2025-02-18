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

// NewCancelInstallationParams creates a new CancelInstallationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelInstallationParams() *CancelInstallationParams {
	return &CancelInstallationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelInstallationParamsWithTimeout creates a new CancelInstallationParams object
// with the ability to set a timeout on a request.
func NewCancelInstallationParamsWithTimeout(timeout time.Duration) *CancelInstallationParams {
	return &CancelInstallationParams{
		timeout: timeout,
	}
}

// NewCancelInstallationParamsWithContext creates a new CancelInstallationParams object
// with the ability to set a context for a request.
func NewCancelInstallationParamsWithContext(ctx context.Context) *CancelInstallationParams {
	return &CancelInstallationParams{
		Context: ctx,
	}
}

// NewCancelInstallationParamsWithHTTPClient creates a new CancelInstallationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelInstallationParamsWithHTTPClient(client *http.Client) *CancelInstallationParams {
	return &CancelInstallationParams{
		HTTPClient: client,
	}
}

/* CancelInstallationParams contains all the parameters to send to the API endpoint
   for the cancel installation operation.

   Typically these are written to a http.Request.
*/
type CancelInstallationParams struct {

	/* ClusterID.

	   The cluster whose installation is to be canceled.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelInstallationParams) WithDefaults() *CancelInstallationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelInstallationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel installation params
func (o *CancelInstallationParams) WithTimeout(timeout time.Duration) *CancelInstallationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel installation params
func (o *CancelInstallationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel installation params
func (o *CancelInstallationParams) WithContext(ctx context.Context) *CancelInstallationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel installation params
func (o *CancelInstallationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel installation params
func (o *CancelInstallationParams) WithHTTPClient(client *http.Client) *CancelInstallationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel installation params
func (o *CancelInstallationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the cancel installation params
func (o *CancelInstallationParams) WithClusterID(clusterID strfmt.UUID) *CancelInstallationParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the cancel installation params
func (o *CancelInstallationParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelInstallationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
