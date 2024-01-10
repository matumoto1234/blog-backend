package infra

import (
	"context"
	"time"

	"github.com/matumoto1234/blog-backend/app/domain/models"
	"github.com/matumoto1234/blog-backend/app/infra/api"
	"github.com/matumoto1234/blog-backend/db"
)

type ArticleRepository interface {
	FindArticles(ctx context.Context, offset, limit int) ([]models.Article, error)
}

var _ ArticleRepository = (*articleRepositoryImpl)(nil)

type articleRepositoryImpl struct {
	client  *db.PrismaClient
	gateway api.Gateway
}

func NewArticleRepository(client *db.PrismaClient, gateway api.Gateway) *articleRepositoryImpl {
	return &articleRepositoryImpl{
		client:  client,
		gateway: gateway,
	}
}

func (r *articleRepositoryImpl) FindArticles(ctx context.Context, offset, limit int) ([]models.Article, error) {
	dtoList, err := r.client.Article.FindMany().Skip(offset).Take(limit).Exec(ctx)
	if err != nil {
		return nil, err
	}

	articles := make([]models.Article, 0, len(dtoList))

	// TODO: goroutineによる並列化
	for _, dto := range dtoList {
		body, err := r.gateway.FetchArticleBody(ctx, dto.ContentURL)
		if err != nil {
			return nil, err
		}

		articles = append(articles, models.Article{
			ID:          dto.ID,
			Title:       dto.Title,
			Body:        string(body),
			PublishedAt: dto.CreatedAt.Format(time.DateOnly),
		})
	}

	return articles, nil
}
