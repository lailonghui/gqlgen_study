//go:generate go run github.com/99designs/gqlgen
package graph

import "lai.com/gqlgen_study/demo01/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PhotoList []*model.Photo
}
