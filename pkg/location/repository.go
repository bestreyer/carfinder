package location

import (
	"context"
)

type Repository interface {
	Update(ctx context.Context, ul *UpdateLocation) (error)
	Create(ctx context.Context, l *Location) (error)
	GetDrivers(ctx context.Context, dal *DriverAroundLocation) ([]LocationWithDistance, error)
}

