package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"lai.com/gqlgen_study/keyVehicleSupervision02/graph/generated"
	"lai.com/gqlgen_study/keyVehicleSupervision02/graph/model"
)

func (r *queryResolver) GetVehicleAlarmDataList(ctx context.Context, paging *model.DataPage, sorting *model.SortDirection) ([]*model.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) NewVehicleAlarmData(ctx context.Context, typeArg *model.AlarmType) (<-chan *model.VehicleAlarmData, error) {
	panic(fmt.Errorf("not implemented"))
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }