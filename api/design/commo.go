package design

import . "goa.design/goa/v3/dsl"

var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a resource that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("user 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing resource")
	Required("message", "id")
})
