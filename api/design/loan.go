package design

import . "goa.design/goa/v3/dsl"

var _ = Service("loan", func() {
	Description("The loan service makes it possible to view, add or remove loans")

	HTTP(func() {
		Path("/loans")
	})

	Method("listLoans", func() {
		Description("List all stored loans")
		Result(CollectionOf(Loan), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("getLoan", func() {
		Description("get loan by ID")
		Payload(func() {
			Field(1, "id", Int, "ID of loan to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
				Default("default")
			})
			Required("id")
		})
		Result(Loan)
		Error("not_found", NotFound, "Loan not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})
})

var Loan = ResultType("application/vnd.cellar.loan", func() {
	Description("A Loan describes a loan in loan service.")
	TypeName("Loan")

	Attributes(func() {
		Attribute("id", Int, "ID is the unique id of the loan", func() {
			Example(1)
		})
		Attribute("amount", Float64, "Loan amount", func() {
			Example(50)
			Minimum(50)
			Maximum(500)
		})
		Attribute("lender_id", Int, "id of lender", func() {
			Example(1)
		})
		Attribute("borrower_id", Int, "id of borrower", func() {
			Example(1)
		})
		Attribute("description", String, "description")
		Attribute("payback_date", String, "Date on which the loan is supposed to be paid back")
	})

	View("default", func() {
		Attribute("id")
		Attribute("amount")
		Attribute("lender_id")
		Attribute("borrower_id")
		Attribute("description")
		Attribute("payback_date")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("amount")
		Attribute("lender_id")
		Attribute("borrower_id")
	})

	Required("id", "amount", "lender_id", "borrower_id")
})
