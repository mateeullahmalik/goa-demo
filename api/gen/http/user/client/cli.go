// Code generated by goa v3.4.3, DO NOT EDIT.
//
// user HTTP client CLI support package
//
// Command:
// $ goa gen github.com/mateeullahmalik/goa-demo/api/design

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	user "github.com/mateeullahmalik/goa-demo/api/gen/user"
	goa "goa.design/goa/v3/pkg"
)

// BuildGetUserPayload builds the payload for the user getUser endpoint from
// CLI flags.
func BuildGetUserPayload(userGetUserID string, userGetUserView string) (*user.GetUserPayload, error) {
	var err error
	var id string
	{
		id = userGetUserID
	}
	var view *string
	{
		if userGetUserView != "" {
			view = &userGetUserView
			if view != nil {
				if !(*view == "default" || *view == "tiny") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &user.GetUserPayload{}
	v.ID = id
	v.View = view

	return v, nil
}

// BuildCreateUserPayload builds the payload for the user createUser endpoint
// from CLI flags.
func BuildCreateUserPayload(userCreateUserBody string) (*user.User, error) {
	var err error
	var body CreateUserRequestBody
	{
		err = json.Unmarshal([]byte(userCreateUserBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"jhon@doe.com\",\n      \"id\": 1,\n      \"name\": \"Jesse owens\",\n      \"score\": 100\n   }'")
		}
		if utf8.RuneCountInString(body.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 3, true))
		}
		if utf8.RuneCountInString(body.Name) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 128, false))
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
		if err != nil {
			return nil, err
		}
	}
	v := &user.User{
		ID:    body.ID,
		Name:  body.Name,
		Email: body.Email,
		Score: body.Score,
	}

	return v, nil
}

// BuildRemoveUserPayload builds the payload for the user removeUser endpoint
// from CLI flags.
func BuildRemoveUserPayload(userRemoveUserID string) (*user.RemoveUserPayload, error) {
	var id string
	{
		id = userRemoveUserID
	}
	v := &user.RemoveUserPayload{}
	v.ID = id

	return v, nil
}
