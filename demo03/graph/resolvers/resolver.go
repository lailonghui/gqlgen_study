//go:generate go run github.com/99designs/gqlgen
package resolvers

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"lai.com/gqlgen_study/demo02/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MutationResolver struct {
		SingleUpload              func(ctx context.Context, file graphql.Upload) (*model.File, error)
		SingleUploadWithPayload   func(ctx context.Context, req model.UploadFile) (*model.File, error)
		MultipleUpload            func(ctx context.Context, files []*graphql.Upload) ([]*model.File, error)
		MultipleUploadWithPayload func(ctx context.Context, req []*model.UploadFile) ([]*model.File, error)
	}
	QueryResolver struct {
		Empty func(ctx context.Context) (string, error)
	}
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}
