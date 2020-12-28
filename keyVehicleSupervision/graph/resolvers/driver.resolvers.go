package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	model1 "keyVehicleSupervision/graph/model"
)

func (r *mutationResolver) CreateDriverInfo(ctx context.Context, req model1.NewDriverInfo) (*model1.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriverInfo(ctx context.Context, req model1.NewDriverInfo) (*model1.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDriverInfo(ctx context.Context, id int) (*model1.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateVehicleDriverBinding(ctx context.Context, req model1.NewVehicleDriverBinding) (*model1.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleDriverBinding(ctx context.Context, req model1.NewVehicleDriverBinding) (*model1.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleDriverBinding(ctx context.Context, id int) (*model1.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDistrictDriverList(ctx context.Context, paging *model1.DataPage, sorting *model1.SortDirection) ([]*model1.DistrictCount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDriverInfoList(ctx context.Context, filter *model1.DriverInfoFilter, paging *model1.DataPage, sorting *model1.DriverDataSort) ([]*model1.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetVehicleDriverBinding(ctx context.Context, filter *model1.VehicleDriverBindingFilter, paging *model1.DataPage, sorting *model1.VehicleDriverBindingDataSort) ([]*model1.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}
