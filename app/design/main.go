package design

import (
	"goa.design/goa/v3/dsl"
)

var _ = dsl.API("blog", func() {
	dsl.Title("Blog Backend")
	dsl.Description("Backend API service for matumoto1234's blog")
	dsl.Server("blog", func() {
		dsl.Host("localhost", func() {
			dsl.URI("http://localhost:8080")
			// URI("grpc://localhost:8080")
		})
	})
})
