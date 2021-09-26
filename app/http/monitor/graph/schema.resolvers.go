package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/si3nloong/webhook/app/http/monitor/graph/generated"
	"github.com/si3nloong/webhook/app/http/monitor/graph/model"
	"github.com/si3nloong/webhook/app/http/monitor/transformer"
)

func (r *queryResolver) Logs(ctx context.Context, after *string, before *string, first *uint, last *uint) (*model.LogConnection, error) {
	limit := uint(100)
	if first != nil {
		limit = *first
	} else if last != nil {
		limit = *last
	}

	datas, _, err := r.GetLogs(ctx, "", limit)
	if err != nil {
		return nil, err
	}

	return transformer.ToLogConnection(datas), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
