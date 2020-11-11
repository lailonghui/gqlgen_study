package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"lai.com/gqlgen_study/keyVehicleSupervision/graph/generated"
	"lai.com/gqlgen_study/keyVehicleSupervision/graph/model"
)

func (r *mutationResolver) CreateVehicleInfo(ctx context.Context, req model.NewVehicleInfo) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfo(ctx context.Context, req model.NewVehicleInfo) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleInfo(ctx context.Context, req model.NewVehicleInfo) (*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDistrictVehicleList(ctx context.Context, paging *model.DataPage, sorting *model.SortDirection) ([]*model.DistrictCount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetVehicleInfoList(ctx context.Context, filter *model.VehicleInfoFilter, paging *model.DataPage, sorting *model.VehicleDataSort) ([]*model.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
