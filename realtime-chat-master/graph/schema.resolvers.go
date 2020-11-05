package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"lai.com/gqlgen_study/graph/generated"
	"lai.com/gqlgen_study/graph/model"
)

func (r *mutationResolver) Post(ctx context.Context, text string, username string, roomName string) (*model.Message, error) {
	r.mu.Lock()
	room := r.Rooms[roomName]
	if room == nil {
		r.Rooms = make(map[string]*model.Chatroom, 10)
		room = &model.Chatroom{
			Name: roomName,
			Observers: map[string]struct {
				Username string
				Message  chan *model.Message
			}{},
		}

	}
	r.mu.Unlock()
	message := model.Message{
		ID:        randString(8),
		CreatedAt: time.Now(),
		Text:      text,
		CreatedBy: username,
	}
	room.Messages = append(room.Messages, message)
	r.Rooms[roomName] = room
	r.mu.Lock()
	for _, observer := range room.Observers {
		if observer.Username == "" || observer.Username == message.CreatedBy {
			observer.Message <- &message
		}
	}
	r.mu.Unlock()
	return &message, nil
}

func (r *queryResolver) Room(ctx context.Context, name string) (*model.Chatroom, error) {
	r.mu.Lock()
	room := r.Rooms[name]
	if room == nil {
		r.Rooms = make(map[string]*model.Chatroom)
		room = &model.Chatroom{
			Name: name,
			Observers: map[string]struct {
				Username string
				Message  chan *model.Message
			}{},
		}
		r.Rooms[name] = room
	}
	r.mu.Unlock()

	return room, nil
}

func (r *subscriptionResolver) MessageAdded(ctx context.Context, roomName string) (<-chan *model.Message, error) {
	fmt.Println("MessageAdded")
	r.mu.Lock()
	room := r.Rooms[roomName]
	if room == nil {
		r.Rooms = make(map[string]*model.Chatroom)
		room = &model.Chatroom{
			Name: roomName,
			Observers: map[string]struct {
				Username string
				Message  chan *model.Message
			}{},
		}
		r.Rooms[roomName] = room
	}
	r.mu.Unlock()

	id := randString(8)
	events := make(chan *model.Message, 1)

	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(room.Observers, id)
		r.mu.Unlock()
	}()

	r.mu.Lock()
	room.Observers[id] = struct {
		Username string
		Message  chan *model.Message
	}{Username: getUsername(ctx), Message: events}
	r.mu.Unlock()

	return events, nil
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
