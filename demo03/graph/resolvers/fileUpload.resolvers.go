package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"lai.com/gqlgen_study/demo02/graph/model"
)

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	return r.MutationResolver.SingleUpload(ctx, file)
}

func (r *mutationResolver) SingleUploadWithPayload(ctx context.Context, req model.UploadFile) (*model.File, error) {
	return r.MutationResolver.SingleUploadWithPayload(ctx, req)
}

func (r *mutationResolver) MultipleUpload(ctx context.Context, files []*graphql.Upload) ([]*model.File, error) {
	return r.MutationResolver.MultipleUpload(ctx, files)
}

func (r *mutationResolver) MultipleUploadWithPayload(ctx context.Context, req []*model.UploadFile) ([]*model.File, error) {
	return r.MutationResolver.MultipleUploadWithPayload(ctx, req)
}

func (r *queryResolver) Empty(ctx context.Context) (string, error) {
	return "hello, lai", nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) MutipleUpload(ctx context.Context, files []*graphql.Upload) ([]*model.File, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) MutipleUploadWithPayload(ctx context.Context, req []*model.UploadFile) (*model.File, error) {
	panic(fmt.Errorf("not implemented"))
}
