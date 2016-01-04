package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewListTasksParams creates a new ListTasksParams object
// with the default values initialized.
func NewListTasksParams() *ListTasksParams {
	return &ListTasksParams{
		PageSize: 20,
	}
}

/*ListTasksParams contains all the parameters to send to the API endpoint
for the list tasks operation typically these are written to a http.Request
*/
type ListTasksParams struct {

	/*PageSize
	  Amount of items to return in a single page

	*/
	PageSize int32
	/*SinceID
	  The last id that was seen.

	*/
	SinceID int64
	/*Status
	  the status to filter by

	*/
	Status []string
	/*Tags
	  the tags to filter by

	*/
	Tags []string
}

// WithPageSize adds the pageSize to the list tasks params
func (o *ListTasksParams) WithPageSize(pageSize int32) *ListTasksParams {
	o.PageSize = pageSize
	return o
}

// WithSinceID adds the sinceId to the list tasks params
func (o *ListTasksParams) WithSinceID(sinceId int64) *ListTasksParams {
	o.SinceID = sinceId
	return o
}

// WithStatus adds the status to the list tasks params
func (o *ListTasksParams) WithStatus(status []string) *ListTasksParams {
	o.Status = status
	return o
}

// WithTags adds the tags to the list tasks params
func (o *ListTasksParams) WithTags(tags []string) *ListTasksParams {
	o.Tags = tags
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListTasksParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// query param pageSize
	qrPageSize := o.PageSize
	qPageSize := swag.FormatInt32(qrPageSize)
	if qPageSize != "" {
		if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
			return err
		}
	}

	// query param sinceId
	qrSinceID := o.SinceID
	qSinceID := swag.FormatInt64(qrSinceID)
	if qSinceID != "" {
		if err := r.SetQueryParam("sinceId", qSinceID); err != nil {
			return err
		}
	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "pipes")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	valuesTags := o.Tags

	joinedTags := swag.JoinByFormat(valuesTags, "")
	// query array param tags
	if err := r.SetQueryParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}