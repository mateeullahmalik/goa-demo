package services

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/mateeullahmalik/goa-demo/api"
	"github.com/mateeullahmalik/goa-demo/api/gen"
	"github.com/mateeullahmalik/goa-demo/api/gen/http/swagger/server"

	goahttp "goa.design/goa/v3/http"
)

// Swagger represents services for swagger endpoints.
type Swagger struct {
	*Common
}

// Mount configures the mux to serve the swagger endpoints.
func (service *Swagger) Mount(_ context.Context, mux goahttp.Muxer) goahttp.Server {
	srv := server.New(nil, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, api.ErrorHandler, nil, nil)

	for _, m := range srv.Mounts {
		file, err := gen.OpenAPIContent.Open(m.Method)
		if err != nil {
			continue
		}

		mux.Handle(m.Verb, m.Pattern, func(w http.ResponseWriter, r *http.Request) {
			http.ServeContent(w, r, m.Method, time.Time{}, file.(io.ReadSeeker))
		})
	}
	return srv
}

// NewSwagger returns the swagger Swagger implementation.
func NewSwagger() *Swagger {
	return &Swagger{
		Common: NewCommon(),
	}
}
