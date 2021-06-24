// Code generated by goa v3.4.3, DO NOT EDIT.
//
// user HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/mateeullahmalik/goa-demo/api/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	user "github.com/mateeullahmalik/goa-demo/api/gen/user"
	userviews "github.com/mateeullahmalik/goa-demo/api/gen/user/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildListUsersRequest instantiates a HTTP request object with method and
// path set to call the "user" service "listUsers" endpoint
func (c *Client) BuildListUsersRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListUsersUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "listUsers", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListUsersResponse returns a decoder for responses returned by the user
// listUsers endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListUsersResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListUsersResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "listUsers", err)
			}
			p := NewListUsersUserCollectionOK(body)
			view := "tiny"
			vres := userviews.UserCollection{Projected: p, View: view}
			if err = userviews.ValidateUserCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "listUsers", err)
			}
			res := user.NewUserCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "listUsers", resp.StatusCode, string(body))
		}
	}
}

// BuildGetUserRequest instantiates a HTTP request object with method and path
// set to call the "user" service "getUser" endpoint
func (c *Client) BuildGetUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*user.GetUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("user", "getUser", "*user.GetUserPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserUserPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "getUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetUserRequest returns an encoder for requests sent to the user
// getUser server.
func EncodeGetUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.GetUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("user", "getUser", "*user.GetUserPayload", v)
		}
		values := req.URL.Query()
		if p.View != nil {
			values.Add("view", *p.View)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeGetUserResponse returns a decoder for responses returned by the user
// getUser endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetUserResponse may return the following errors:
//	- "not_found" (type *user.NotFound): http.StatusNotFound
//	- error: internal error
func DecodeGetUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "getUser", err)
			}
			p := NewGetUserUserOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &userviews.User{Projected: p, View: view}
			if err = userviews.ValidateUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "getUser", err)
			}
			res := user.NewUser(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetUserNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "getUser", err)
			}
			err = ValidateGetUserNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "getUser", err)
			}
			return nil, NewGetUserNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "getUser", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateUserRequest instantiates a HTTP request object with method and
// path set to call the "user" service "createUser" endpoint
func (c *Client) BuildCreateUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateUserUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "createUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateUserRequest returns an encoder for requests sent to the user
// createUser server.
func EncodeCreateUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.User)
		if !ok {
			return goahttp.ErrInvalidType("user", "createUser", "*user.User", v)
		}
		body := NewCreateUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "createUser", err)
		}
		return nil
	}
}

// DecodeCreateUserResponse returns a decoder for responses returned by the
// user createUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCreateUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "createUser", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "createUser", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveUserRequest instantiates a HTTP request object with method and
// path set to call the "user" service "removeUser" endpoint
func (c *Client) BuildRemoveUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*user.RemoveUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("user", "removeUser", "*user.RemoveUserPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemoveUserUserPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "removeUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRemoveUserResponse returns a decoder for responses returned by the
// user removeUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeRemoveUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "removeUser", resp.StatusCode, string(body))
		}
	}
}

// unmarshalUserResponseToUserviewsUserView builds a value of type
// *userviews.UserView from a value of type *UserResponse.
func unmarshalUserResponseToUserviewsUserView(v *UserResponse) *userviews.UserView {
	res := &userviews.UserView{
		ID:    v.ID,
		Name:  v.Name,
		Email: v.Email,
		Score: v.Score,
	}

	return res
}
