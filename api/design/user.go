package design

import . "goa.design/goa/v3/dsl"

var _ = Service("user", func() {
	Description("The user service makes it possible to view, add or remove users")

	HTTP(func() {
		Path("/users")
	})

	Method("listUsers", func() {
		Description("List all stored users")
		Result(CollectionOf(User), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("getUser", func() {
		Description("get user by ID")
		Payload(func() {
			Field(1, "id", String, "ID of user to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(User)
		Error("not_found", NotFound, "User not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("createUser", func() {
		Description("create new user and return its ID.")
		Payload(User)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("removeUser", func() {
		Description("Remove user")
		Payload(func() {
			Field(1, "id", String, "ID of user to remove")
			Required("id")
		})
		Error("not_found", NotFound, "User not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})

var User = ResultType("application/vnd.cellar.user", func() {
	Description("A User describes a user retrieved by the user service.")
	TypeName("User")

	Attributes(func() {
		Attribute("id", Int, "ID is the unique id of the user.", func() {
			Example(1)
		})
		Attribute("name", String, "Full Name of the user ", func() {
			Example("Jesse owens")
			MinLength(3)
			MaxLength(128)
		})
		Attribute("email", String, "email address of user", func() {
			Example("jhon@doe.com")
		})
		Attribute("score", Int, "Score is user's loan score", func() {
			Example(100)
			Maximum(1000)
			Minimum(0)
		})
	})

	View("default", func() {
		Attribute("id")
		Attribute("score")
		Attribute("name")
		Attribute("email")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
	})

	Required("id", "name", "email")
})
