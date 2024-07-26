package service_domains

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type PlacementServiceInterface interface {
	CreatePlacement(ctx context.Context, data forms.PlacementCreateForm) (models.Placement, error)
	GetPlacementOrder(ctx context.Context, id uint) (models.Placement, error)
	Find(ctx context.Context, id uint) (models.Placement, error)
	Read(ctx context.Context, q string, page uint) ([]models.Placement, *types.Pagination, error)
	Update(ctx context.Context, id uint, data forms.PlacementUpdateForm) (uint, models.Placement, error)
	Delete(ctx context.Context, id uint) (uint, error)
	PlacementDetails(ctx context.Context, id uint) ([]models.PlacementDetail, error)
	PlaceNewEmployee(ctx context.Context, data forms.PlacementDetailEmployeeCreateForm) (models.PlacementDetailEmployee, error)
	EmployeePlacementDetail(ctx context.Context, id uint) (models.PlacementDetailEmployee, error)
	CutoffEmployeePlacement(ctx context.Context, placementDetailEmployeeID uint, data forms.PlacementDetailEmployeeUpdateForm) (uint, error)
	SuspendEmployeePlacement(ctx context.Context, placementDetailEmployeeID uint) (uint, error)
	OngoingEmployeePlacement(ctx context.Context, placementDetailEmployeeID uint) (uint, error)
	RemoveEmployeePlacement(ctx context.Context, placementDetailEmployeeID uint) (uint, error)
	SetDoneEmployeePlacement(ctx context.Context, placementDetailEmployeeID uint) error
}
