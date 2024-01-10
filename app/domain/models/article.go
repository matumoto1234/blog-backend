package models

type Article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	PublishedAt string `json:"published_at"`
}

type ArticleSummary struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	SummaryBody string `json:"summary_body"`
	PublishedAt string `json:"published_at"`
}
