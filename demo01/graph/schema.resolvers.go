package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"lai.com/gqlgen_study/demo01/graph/generated"
	"lai.com/gqlgen_study/demo01/graph/model"
)

func (r *mutationResolver) PostPhoto(ctx context.Context, input model.PostPhotoInput) (*model.Photo, error) {
	p := &model.Photo{
		ID:          strconv.Itoa(num),
		Name:        input.Name,
		Description: input.Description,
		Category:    *input.Category,
		Created:     time.Now(),
	}
	r.PhotoList = append(r.PhotoList, p)
	return p, nil
}

func (r *mutationResolver) GithubAuth(ctx context.Context, code *string) (*model.AuthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TotalPhoto(ctx context.Context) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllPhotos(ctx context.Context, filter *model.PhotoFilter, paging *model.DataPage, sorting *model.DataSort) ([]*model.Photo, error) {
	return r.PhotoList, nil
}

func (r *queryResolver) Photo(ctx context.Context, id string) (*model.Photo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TotalUser(ctx context.Context) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllUsers(ctx context.Context, paging *model.DataPage, sorting *model.DataSort) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, githubLogin string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Agenda(ctx context.Context) ([]model.AgendaItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) NewPhoto(ctx context.Context, category *model.PhotoCategory) (<-chan *model.Photo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) NewUser(ctx context.Context) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var num = 1
