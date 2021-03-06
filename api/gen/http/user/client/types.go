// Code generated by goa v3.4.3, DO NOT EDIT.
//
// user HTTP client types
//
// Command:
// $ goa gen github.com/mateeullahmalik/goa-demo/api/design

package client

import (
	"unicode/utf8"

	user "github.com/mateeullahmalik/goa-demo/api/gen/user"
	userviews "github.com/mateeullahmalik/goa-demo/api/gen/user/views"
	goa "goa.design/goa/v3/pkg"
)

// CreateUserRequestBody is the type of the "user" service "createUser"
// endpoint HTTP request body.
type CreateUserRequestBody struct {
	// ID is the unique id of the user.
	ID int `form:"id" json:"id" xml:"id"`
	// Full Name of the user
	Name string `form:"name" json:"name" xml:"name"`
	// email address of user
	Email string `form:"email" json:"email" xml:"email"`
	// Score is user's loan score
	Score *int `form:"score,omitempty" json:"score,omitempty" xml:"score,omitempty"`
}

// ListUsersResponseBody is the type of the "user" service "listUsers" endpoint
// HTTP response body.
type ListUsersResponseBody []*UserResponse

// GetUserResponseBody is the type of the "user" service "getUser" endpoint
// HTTP response body.
type GetUserResponseBody struct {
	// ID is the unique id of the user.
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Full Name of the user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// email address of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Score is user's loan score
	Score *int `form:"score,omitempty" json:"score,omitempty" xml:"score,omitempty"`
}

// GetUserNotFoundResponseBody is the type of the "user" service "getUser"
// endpoint HTTP response body for the "not_found" error.
type GetUserNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// ID of missing resource
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// UserResponse is used to define fields on response body types.
type UserResponse struct {
	// ID is the unique id of the user.
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Full Name of the user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// email address of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Score is user's loan score
	Score *int `form:"score,omitempty" json:"score,omitempty" xml:"score,omitempty"`
}

// NewCreateUserRequestBody builds the HTTP request body from the payload of
// the "createUser" endpoint of the "user" service.
func NewCreateUserRequestBody(p *user.User) *CreateUserRequestBody {
	body := &CreateUserRequestBody{
		ID:    p.ID,
		Name:  p.Name,
		Email: p.Email,
		Score: p.Score,
	}
	return body
}

// NewListUsersUserCollectionOK builds a "user" service "listUsers" endpoint
// result from a HTTP "OK" response.
func NewListUsersUserCollectionOK(body ListUsersResponseBody) userviews.UserCollectionView {
	v := make([]*userviews.UserView, len(body))
	for i, val := range body {
		v[i] = unmarshalUserResponseToUserviewsUserView(val)
	}

	return v
}

// NewGetUserUserOK builds a "user" service "getUser" endpoint result from a
// HTTP "OK" response.
func NewGetUserUserOK(body *GetUserResponseBody) *userviews.UserView {
	v := &userviews.UserView{
		ID:    body.ID,
		Name:  body.Name,
		Email: body.Email,
		Score: body.Score,
	}

	return v
}

// NewGetUserNotFound builds a user service getUser endpoint not_found error.
func NewGetUserNotFound(body *GetUserNotFoundResponseBody) *user.NotFound {
	v := &user.NotFound{
		Message: *body.Message,
		ID:      *body.ID,
	}

	return v
}

// ValidateGetUserNotFoundResponseBody runs the validations defined on
// getUser_not_found_response_body
func ValidateGetUserNotFoundResponseBody(body *GetUserNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateUserResponse runs the validations defined on UserResponse
func ValidateUserResponse(body *UserResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 3, true))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 128, false))
		}
	}
	if body.Score != nil {
		if *body.Score < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.score", *body.Score, 0, true))
		}
	}
	if body.Score != nil {
		if *body.Score > 1000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.score", *body.Score, 1000, false))
		}
	}
	return
}
