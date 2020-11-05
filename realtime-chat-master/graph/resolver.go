//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"context"
	"lai.com/gqlgen_study/graph/model"
	"math/rand"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type ckey string
type Resolver struct {
	Rooms map[string]*model.Chatroom
	mu    sync.Mutex // nolint: structcheck
}

func getUsername(ctx context.Context) string {
	if username, ok := ctx.Value(ckey("username")).(string); ok {
		return username
	}
	return ""
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
