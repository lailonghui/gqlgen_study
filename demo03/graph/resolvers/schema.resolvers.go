package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"lai.com/gqlgen_study/demo03/graph/generated"
	"lai.com/gqlgen_study/demo03/graph/model"
	"lai.com/gqlgen_study/demo03/graph/mypkg"
)

func (r *carResolver) Y(ctx context.Context, obj *model.Car) (mypkg.YesNo, error) {
	//y := mypkg.YesNo(true)
	return false, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (*bool, error) {
	return nil, nil
}

func (r *queryResolver) Car(ctx context.Context) (*model.Car, error) {
	c := &model.Car{
		Make:  "11",
		Model: "22",
		Color: "33",
	}
	return c, nil
	//panic(fmt.Errorf("not implemented"))
}

// Car returns generated.CarResolver implementation.
func (r *Resolver) Car() generated.CarResolver { return &carResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type carResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
