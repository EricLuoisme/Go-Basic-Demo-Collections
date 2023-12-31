package serializer

import (
	"context"
	"fmt"
	repository2 "own/example/bolbol/demo/repository"
	"strings"
)

// SimpleSummaryArticle is a JSON format struct
type SimpleSummaryArticle struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	More    string `json:"more"`
}

type Article struct {
	articles          repository2.ArticleRepository
	summaryWordsLimit int
}

// NewArticle is for create serializer struct of Article
func NewArticle(articles repository2.ArticleRepository, summaryWordsLimit int) *Article {
	return &Article{articles: articles, summaryWordsLimit: summaryWordsLimit}
}

// ByID is find by Id
func (a *Article) ByID(ctx context.Context, id uint64) (SimpleSummaryArticle, error) {
	article, err := a.articles.ByID(ctx, int(id))
	if err != nil {
		return SimpleSummaryArticle{}, repository2.ErrNotFound
	}
	return SimpleSummaryArticle{
		ID:      article.ID,
		Title:   article.Title,
		Summary: a.summarize(article.Content),
		More:    fmt.Sprintf("https://site.com/a/%d", article.ID),
	}, nil
}

// summarize is for summarize the content
func (a *Article) summarize(content string) string {
	words := strings.Split(strings.ReplaceAll(content, "\n", " "), " ")
	if len(words) > a.summaryWordsLimit {
		words = words[:a.summaryWordsLimit]
	}
	return strings.Join(words, " ")
}
