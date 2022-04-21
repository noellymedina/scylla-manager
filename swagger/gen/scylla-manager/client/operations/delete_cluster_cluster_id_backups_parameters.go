// Code generated by go-swagger; DO NOT EDIT.

package operations

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
	"github.com/go-openapi/swag"
)

// NewDeleteClusterClusterIDBackupsParams creates a new DeleteClusterClusterIDBackupsParams object
// with the default values initialized.
func NewDeleteClusterClusterIDBackupsParams() *DeleteClusterClusterIDBackupsParams {
	var ()
	return &DeleteClusterClusterIDBackupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterClusterIDBackupsParamsWithTimeout creates a new DeleteClusterClusterIDBackupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterClusterIDBackupsParamsWithTimeout(timeout time.Duration) *DeleteClusterClusterIDBackupsParams {
	var ()
	return &DeleteClusterClusterIDBackupsParams{

		timeout: timeout,
	}
}

// NewDeleteClusterClusterIDBackupsParamsWithContext creates a new DeleteClusterClusterIDBackupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterClusterIDBackupsParamsWithContext(ctx context.Context) *DeleteClusterClusterIDBackupsParams {
	var ()
	return &DeleteClusterClusterIDBackupsParams{

		Context: ctx,
	}
}

// NewDeleteClusterClusterIDBackupsParamsWithHTTPClient creates a new DeleteClusterClusterIDBackupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterClusterIDBackupsParamsWithHTTPClient(client *http.Client) *DeleteClusterClusterIDBackupsParams {
	var ()
	return &DeleteClusterClusterIDBackupsParams{
		HTTPClient: client,
	}
}

/*DeleteClusterClusterIDBackupsParams contains all the parameters to send to the API endpoint
for the delete cluster cluster ID backups operation typically these are written to a http.Request
*/
type DeleteClusterClusterIDBackupsParams struct {

	/*ClusterID*/
	ClusterID string
	/*Locations*/
	Locations []string
	/*ManifestParallelism*/
	ManifestParallelism int64
	/*SnapshotTags*/
	SnapshotTags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) WithTimeout(timeout time.Duration) *DeleteClusterClusterIDBackupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) WithContext(ctx context.Context) *DeleteClusterClusterIDBackupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) WithHTTPClient(client *http.Client) *DeleteClusterClusterIDBackupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) WithClusterID(clusterID string) *DeleteClusterClusterIDBackupsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocations adds the locations to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) WithLocations(locations []string) *DeleteClusterClusterIDBackupsParams {
	o.SetLocations(locations)
	return o
}

// SetLocations adds the locations to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) SetLocations(locations []string) {
	o.Locations = locations
}

// WithManifestParallelism adds the manifestParallelism to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) WithManifestParallelism(manifestParallelism int64) *DeleteClusterClusterIDBackupsParams {
	o.SetManifestParallelism(manifestParallelism)
	return o
}

// SetManifestParallelism adds the manifestParallelism to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) SetManifestParallelism(manifestParallelism int64) {
	o.ManifestParallelism = manifestParallelism
}

// WithSnapshotTags adds the snapshotTags to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) WithSnapshotTags(snapshotTags []string) *DeleteClusterClusterIDBackupsParams {
	o.SetSnapshotTags(snapshotTags)
	return o
}

// SetSnapshotTags adds the snapshotTags to the delete cluster cluster ID backups params
func (o *DeleteClusterClusterIDBackupsParams) SetSnapshotTags(snapshotTags []string) {
	o.SnapshotTags = snapshotTags
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterClusterIDBackupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	valuesLocations := o.Locations

	joinedLocations := swag.JoinByFormat(valuesLocations, "")
	// query array param locations
	if err := r.SetQueryParam("locations", joinedLocations...); err != nil {
		return err
	}

	// query param manifest_parallelism
	qrManifestParallelism := o.ManifestParallelism
	qManifestParallelism := swag.FormatInt64(qrManifestParallelism)
	if qManifestParallelism != "" {
		if err := r.SetQueryParam("manifest_parallelism", qManifestParallelism); err != nil {
			return err
		}
	}

	valuesSnapshotTags := o.SnapshotTags

	joinedSnapshotTags := swag.JoinByFormat(valuesSnapshotTags, "")
	// query array param snapshot_tags
	if err := r.SetQueryParam("snapshot_tags", joinedSnapshotTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
