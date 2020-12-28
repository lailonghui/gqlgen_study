package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	generated1 "keyVehicleSupervision/graph/generated"
	model1 "keyVehicleSupervision/graph/model"
)

func (r *mutationResolver) CreateVehicleInfo(ctx context.Context, req model1.NewVehicleInfo) (*model1.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleInfo(ctx context.Context, req model1.NewVehicleInfo) (*model1.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleInfo(ctx context.Context, id int) (*model1.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDistrictVehicleList(ctx context.Context, paging *model1.DataPage, sorting *model1.SortDirection) ([]*model1.DistrictCount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetVehicleInfoList(ctx context.Context, filter *model1.VehicleInfoFilter, paging *model1.DataPage, sorting *model1.VehicleDataSort) ([]*model1.VehicleInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
