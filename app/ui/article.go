package ui

import (
	"context"

	"github.com/matumoto1234/blog-backend/app/middleware/applog"
	"github.com/matumoto1234/blog-backend/app/usecase"
	"github.com/matumoto1234/blog-backend/gen/article"
)

var _ article.Service = (*articleHandler)(nil)

type articleHandler struct {
	u usecase.ArticleUseCase
}

func NewArticleHandler(u usecase.ArticleUseCase) *articleHandler {
	return &articleHandler{
		u: u,
	}
}

func (h *articleHandler) Summaries(ctx context.Context, p *article.SummariesPayload) (*article.ArticleSummaryResponse, error) {
	ctx = applog.WithContext(ctx, applog.New())

	applog.FromContext(ctx).Info("received request: Summaries")

	summaries, err := h.u.Summaries(ctx, int(p.Offset), int(p.Limit))
	if err != nil {
		return nil, err
	}

	resSummaries := make([]*article.ArticleSummary, len(summaries))

	for i := range summaries {
		resSummaries[i] = &article.ArticleSummary{
			ArticleID:   summaries[i].ID,
			Title:       summaries[i].Title,
			SummaryBody: summaries[i].SummaryBody,
			PublishedAt: summaries[i].PublishedAt,
		}
	}

	return &article.ArticleSummaryResponse{
		ArticleSummaries: resSummaries,
		TotalCount:       uint(len(resSummaries)),
	}, nil
}
