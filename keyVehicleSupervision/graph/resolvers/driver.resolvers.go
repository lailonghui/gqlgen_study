package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"lai.com/gqlgen_study/keyVehicleSupervision/graph/model"
)

func (r *mutationResolver) CreateDriverInfo(ctx context.Context, req model.NewDriverInfo) (*model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDriverInfo(ctx context.Context, req model.NewDriverInfo) (*model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDriverInfo(ctx context.Context, req model.NewDriverInfo) (*model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateVehicleDriverBinding(ctx context.Context, req model.NewVehicleDriverBinding) (*model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVehicleDriverBinding(ctx context.Context, req model.NewVehicleDriverBinding) (*model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicleDriverBinding(ctx context.Context, req model.NewVehicleDriverBinding) (*model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDistrictDriverList(ctx context.Context, paging *model.DataPage, sorting *model.SortDirection) ([]*model.DistrictCount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDriverInfoList(ctx context.Context, filter *model.DriverInfoFilter, paging *model.DataPage, sorting *model.DriverDataSort) ([]*model.DriverInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetVehicleDriverBinding(ctx context.Context, filter *model.DriverInfoFilter, paging *model.DataPage, sorting *model.DriverDataSort) ([]*model.VehicleDriverBinding, error) {
	panic(fmt.Errorf("not implemented"))
}
