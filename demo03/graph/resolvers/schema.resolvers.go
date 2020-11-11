package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"lai.com/gqlgen_study/demo02/graph/generated"
	"lai.com/gqlgen_study/demo02/graph/model"
)

func (r *mutationResolver) Users(ctx context.Context) ([]*model.User, error) {
	return nil, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	//graphql.AddErrorf(ctx, "Error %d", 1)
	//
	//graphql.AddError(ctx, gqlerror.Errorf("zzztt"))

	graphql.AddError(ctx, &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "A descriptive error message",
		Extensions: map[string]interface{}{
			"code": "10-4",
		},
	})

	//return nil, gqlerror.Errorf("BOOM! Headshot")
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
