package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/examples/tutorials/todo-list/second-generation/models"
)

// NewUpdateOneParams creates a new UpdateOneParams object
// with the default values initialized.
func NewUpdateOneParams() UpdateOneParams {
	return UpdateOneParams{}
}

// UpdateOneParams contains all the bound params for the update one operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateOne
type UpdateOneParams struct {
	/*
	  In: body
	*/
	Body *models.Item
	/*
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateOneParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	var body models.Item
	if err := route.Consumer.Consume(r.Body, &body); err != nil {
		res = append(res, errors.NewParseError("body", "body", "", err))
	} else {
		if err := body.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

		if len(res) == 0 {
			o.Body = &body
		}
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOneParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}