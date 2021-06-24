package services

import (
	"context"
	"fmt"

	"github.com/mateeullahmalik/goa-demo/api"
	"github.com/mateeullahmalik/goa-demo/internal/log"

	"github.com/mateeullahmalik/goa-demo/api/gen/http/user/server"
	"github.com/mateeullahmalik/goa-demo/api/gen/user"
	usersrvc "github.com/mateeullahmalik/goa-demo/services/user"

	goahttp "goa.design/goa/v3/http"
)

// UserSrvc represents services for user endpoints.
type UserSrvc struct {
	*Common
	user *usersrvc.Service
	// db
	// other assets
}

// ListUsers returns users
func (service *UserSrvc) ListUsers(context.Context) (res user.UserCollection, err error) {
	fmt.Println("ListUsers called")

	return
}

// get user by ID
// The "view" return value must have one of the following views
//	- "default"
//	- "tiny"
func (service *UserSrvc) GetUser(context.Context, *user.GetUserPayload) (res *user.User, view string, err error) {
	fmt.Println("GetUser called")

	return
}

// create new user and return its ID.
func (service *UserSrvc) CreateUser(context.Context, *user.User) (res string, err error) {
	fmt.Println("CreateUser called")

	return
}

// Remove user
func (service *UserSrvc) RemoveUser(context.Context, *user.RemoveUserPayload) (err error) {
	fmt.Println("RemoveUser called")

	return
}

// Mount configures the mux to serve the user endpoints.
func (service *UserSrvc) Mount(ctx context.Context, mux goahttp.Muxer) goahttp.Server {
	endpoints := user.NewEndpoints(service)
	srv := server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, api.ErrorHandler, nil)
	server.Mount(mux, srv)

	for _, m := range srv.Mounts {
		log.WithContext(ctx).Infof("%q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	return srv
}

// NewUserSrvc returns the users User implementation.
func NewUserSrvc(srvc *usersrvc.Service) *UserSrvc {
	return &UserSrvc{
		Common: NewCommon(),
		user:   srvc,
	}
}
