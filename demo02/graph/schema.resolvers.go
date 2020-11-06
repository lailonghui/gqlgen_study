package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"lai.com/gqlgen_study/demo02/graph/dataloader"

	"lai.com/gqlgen_study/demo02/graph/generated"
	"lai.com/gqlgen_study/demo02/graph/model"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	fmt.Println("todo")
	todos := []*model.Todo{
		{
			ID:     "11",
			Todo:   "study",
			UserID: 1,
		},
		{
			ID:     "22",
			Todo:   "study2",
			UserID: 1,
		},
		{
			ID:     "33",
			Todo:   "study3",
			UserID: 2,
		},
	}
	return todos, nil
	//panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) UserRaw(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) UserLoader(ctx context.Context, obj *model.Todo) (*model.User, error) {
	//fmt.Println("user", obj.UserID)
	//user :=  &model.User{
	//	ID: obj.UserID,
	//	Name: "lai",
	//}
	//
	//return user, nil
	return dataloader.For(ctx).UserById.Load(obj.UserID)

}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
