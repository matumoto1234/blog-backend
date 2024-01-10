package usecase

import (
	"context"

	"github.com/matumoto1234/blog-backend/app/domain/models"
	"github.com/matumoto1234/blog-backend/app/infra"
)

type ArticleUseCase interface {
	Summaries(ctx context.Context, offset, limit int) ([]*models.ArticleSummary, error)
}

var _ ArticleUseCase = (*articleUseCase)(nil)

type articleUseCase struct {
	repo infra.ArticleRepository
}

func NewArticleUseCase(repo infra.ArticleRepository) *articleUseCase {
	return &articleUseCase{
		repo: repo,
	}
}

func (u *articleUseCase) Summaries(ctx context.Context, offset, limit int) ([]*models.ArticleSummary, error) {
	articles, err := u.repo.FindArticles(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	var summaries []*models.ArticleSummary

	// articles[i].Bodyがメモリ的に大きいので、インデックスアクセスで行っている
	for i := range articles {

		// 200文字までのサマリーを生成
		summaryLen := 200

		if len(articles[i].Body) < summaryLen {
			summaryLen = len(articles[i].Body)
		}

		summaries = append(summaries, &models.ArticleSummary{
			ID:          articles[i].ID,
			Title:       articles[i].Title,
			SummaryBody: articles[i].Body[:summaryLen],
			PublishedAt: articles[i].PublishedAt,
		})
	}

	return summaries, nil
}
