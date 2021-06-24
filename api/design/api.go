package design

import . "goa.design/goa/v3/dsl"

var _ = API("demo", func() {
	Title("Demo Service")
	Description("HTTP service for managing user & their loans")

	Server("demo", func() {
		Description("demo hosts the loan, user and swagger services.")
		Services("swagger", "user", "loan")
		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})
