package api

import (
	"context"
	"io"
	"net/http"
)

type Gateway interface {
	FetchArticleBody(ctx context.Context, url string) ([]byte, error)
}

var _ Gateway = (*gatewayImpl)(nil)

type gatewayImpl struct {
	client *http.Client
}

func NewGateway(client *http.Client) *gatewayImpl {
	return &gatewayImpl{
		client: client,
	}
}

func (g *gatewayImpl) FetchArticleBody(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
