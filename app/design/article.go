package design

import (
	"goa.design/goa/v3/dsl"
)

var _ = dsl.Service("article", func() {
	dsl.Description("Article service")

	dsl.Method("summaries", func() {
		dsl.Description("Get article summaries")

		dsl.Payload(func() {
			dsl.Attribute("limit", dsl.UInt, "Limit of article summaries", func() {
				dsl.Default(20)
				dsl.Example(20)
			})
			dsl.Attribute("offset", dsl.UInt, "Offset of article summaries", func() {
				dsl.Default(0)
				dsl.Example(0)
			})
		})

		dsl.Result(ArticleSummariesResponse)

		dsl.Error("BadRequest", CustomError)
		dsl.Error("InternalServerError", CustomError)

		dsl.HTTP(func() {
			dsl.GET("/articles/summaries")
			dsl.Param("limit")
			dsl.Param("offset")
			dsl.Response(dsl.StatusOK)
			dsl.Response("BadRequest", dsl.StatusBadRequest)
			dsl.Response("InternalServerError", dsl.StatusInternalServerError)
		})

		// GRPC(func() {
		// 	Response(CodeOK)
		// })
	})
})

var ArticleSummary = dsl.Type("ArticleSummary", func() {
	dsl.Description("Summary of article")

	dsl.Attribute("article_id", dsl.String, "ID of article", func() {
		dsl.Example("2022hurikaeri")
	})
	dsl.Attribute("title", dsl.String, "Title of article", func() {
		dsl.Example("2022年振り返り")
	})
	dsl.Attribute("summary_body", dsl.String, "Summary of article body", func() {
		dsl.Example("2022年の振り返りをします。")
	})
	dsl.Attribute("published_at", dsl.String, "Published at time of article", func() {
		// format is "2006-01-02".
		// see detail: https://github.com/goadesign/goa/blob/4ce46caf331593edc71dd83cd8ac5641b34ea1de/expr/example.go#L197
		dsl.Format(dsl.FormatDate)
		dsl.Example("2022-01-01")
	})
	dsl.Required("article_id", "title", "summary_body", "published_at")
})

var ArticleSummariesResponse = dsl.Type("ArticleSummaryResponse", func() {
	dsl.Description("Article summaries response")

	dsl.Attribute("article_summaries", dsl.ArrayOf(ArticleSummary), "Article summaries")
	dsl.Attribute("total_count", dsl.UInt, "Total count of article summaries", func() {
		// Exampleとして生成されるarticle_summariesは長さが指定できないため、適当な値をtotalCountに設定しておく。
		dsl.Example(1)
	})

	dsl.Required("article_summaries", "total_count")
})

var CustomError = dsl.Type("CustomError", func() {
	dsl.ErrorName("name", dsl.String, "Name of error.")
	dsl.Attribute("message", dsl.String, "Message of error.")
	dsl.Required("name", "message")
})
