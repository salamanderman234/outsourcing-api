package services

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

type orderService struct{}

func NewOrderService() domains.OrderService {
	return &orderService{}
}

func (orderService) MakeOrder(c context.Context, user domains.UserModel, orderData domains.ServiceOrderForm) (domains.ServiceOrderEntity, error) {
	if ok, err := helpers.Validate(orderData); !ok {
		return domains.ServiceOrderEntity{}, err
	}
	return domains.ServiceOrderEntity{}, nil
}
func (orderService) CancelOrder(c context.Context, user domains.UserModel, orderId uint) (bool, error) {
	cancelStatus := domains.CancelledOrderStatus
	data := domains.ServiceOrderModel{
		Status: &cancelStatus,
	}
	_, _, err := domains.RepoRegistry.ServiceOrderRepo.Update(c, orderId, data)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (orderService) ListOrder(c context.Context,
	user domains.UserModel,
	serviceUserId uint,
	status string,
	page uint,
	orderBy string,
	desc bool,
	withPagination bool) ([]domains.ServiceOrderEntity, *domains.Pagination, error) {

	var pagination domains.Pagination
	orders, maxPage, err := domains.RepoRegistry.ServiceOrderRepo.Read(c,
		string(status), serviceUserId, page, orderBy, desc, withPagination,
	)
	if err != nil {
		return nil, nil, err
	}
	var ordersEntity []domains.ServiceOrderEntity
	for _, order := range orders {
		var orderEntity domains.ServiceOrderEntity
		if err := helpers.Convert(order, &orderEntity); err != nil {
			return nil, nil, domains.ErrConversionType
		}
		ordersEntity = append(ordersEntity, orderEntity)
	}
	if withPagination {
		queries := helpers.MakeDefaultGetPaginationQueries("", 0, page, orderBy, desc, withPagination)
		if serviceUserId != 0 {
			queries["service_user_id"] = serviceUserId
		}
		if status != "" {
			queries["status"] = status
		}
		pagination = helpers.MakePagination(maxPage, uint(page), queries)
		return ordersEntity, &pagination, nil
	}
	return ordersEntity, nil, nil
}
func (orderService) DetailOrder(c context.Context, user domains.UserModel, id uint) (domains.ServiceOrderEntity, error) {
	var order domains.ServiceOrderEntity
	mod, err := domains.RepoRegistry.ServiceOrderRepo.Find(c, id)
	if err != nil {
		return order, err
	}
	if err := helpers.Convert(mod, &order); err != nil {
		return order, err
	}
	return order, err
}
func (orderService) UpdateOrderStatus(c context.Context, user domains.UserModel, id uint, data domains.ServiceOrderUpdateStatusForm) (int, domains.ServiceOrderEntity, error) {
	var dataModel domains.ServiceOrderModel
	var dataEntity domains.ServiceOrderEntity
	fun := func(id uint) (int, domains.Model, error) {
		aff, updated, err := domains.RepoRegistry.ServiceOrderRepo.Update(c, id, dataModel)
		return int(aff), updated, err
	}
	aff, _, err := basicUpdateService(id, data, &dataModel, &dataEntity, fun)
	if err != nil {
		return 0, dataEntity, err
	}
	return aff, dataEntity, nil
}
