package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

type orderService struct{}

func NewOrderService() domains.OrderService {
	return &orderService{}
}

func (orderService) MakeOrder(c context.Context, user domains.UserModel, orderData domains.ServiceOrderForm) (domains.ServiceOrderEntity, error) {
	var orderEntity domains.ServiceOrderEntity
	var orderModel domains.ServiceOrderModel
	details := orderData.Details
	orderData.Details = nil
	if ok, err := helpers.Validate(orderData); !ok {
		return domains.ServiceOrderEntity{}, err
	}
	var orderTotalPrice uint64
	for indexDetail, detail := range details {
		detailItems := detail.Items
		detail.Items = nil
		if ok, err := helpers.Validate(detail); !ok {
			return domains.ServiceOrderEntity{}, err
		}
		psId := detail.PartialServiceID
		ps, err := domains.RepoRegistry.ServiceRepo.Find(c, psId)
		if err != nil {
			conv := domains.ErrForeignKeyViolated
			conv.ValidationErrors = govalidator.Errors{
				govalidator.Error{
					Name:                     fmt.Sprintf("order_details[%d].partial_service_id", indexDetail),
					CustomErrorMessageExists: true,
					Validator:                "foreign key",
					Err:                      errors.New("this partial service does not exists"),
				},
			}
			return orderEntity, conv
		}
		var additionalPrice uint64
		detailData := domains.ServiceOrderDetailModel{
			PartialServiceID: &ps.ID,
			ServicePrice:     ps.BasePrice,
		}
		for indexItem, item := range detailItems {
			if ok, err := helpers.Validate(item); !ok {
				return domains.ServiceOrderEntity{}, err
			}
			psId := detail.PartialServiceID
			si, err := domains.RepoRegistry.ServiceItemRepo.Find(c, psId)
			if err != nil {
				conv := domains.ErrForeignKeyViolated
				conv.ValidationErrors = govalidator.Errors{
					govalidator.Error{
						Name:                     fmt.Sprintf("order_details[%d].partial_service_id.order_detail_items[%d].partial_service_item_id", indexDetail, indexItem),
						CustomErrorMessageExists: true,
						Validator:                "foreign key",
						Err:                      errors.New("this partial service item does not exists"),
					},
				}
				return orderEntity, conv
			}
			if si.Service.ID != ps.ID {
				conv := domains.ErrUnmatchedData
				conv.ValidationErrors = govalidator.Errors{
					govalidator.Error{
						Name:                     fmt.Sprintf("order_details[%d].partial_service_id.order_detail_items[%d].partial_service_item_id", indexDetail, indexItem),
						CustomErrorMessageExists: true,
						Validator:                "unmatched",
						Err:                      errors.New("this item cannot be paired with this partial service"),
					},
				}
				return orderEntity, conv
			}
			total := *si.PricePerItem * uint64(item.Value)
			additionalPrice += total
			detailDataItem := domains.ServiceOrderDetailItemModel{
				Value:                &item.Value,
				PartialServiceItemID: &si.ID,
				ItemPrice:            si.PricePerItem,
				TotalPrice:           &total,
			}
			detailData.PartialServiceItems = append(detailData.PartialServiceItems, detailDataItem)
		}
		detailTotalPrice := additionalPrice + *detailData.ServicePrice
		detailData.AdditionalPrice = &additionalPrice
		detailData.TotalPrice = &detailTotalPrice
		orderTotalPrice += detailTotalPrice
		orderModel.ServiceOrderDetails = append(orderModel.ServiceOrderDetails, detailData)
	}
	finalPrice := orderTotalPrice * uint64(orderData.ContractDuration)
	orderModel.TotalPrice = &finalPrice
	orderModel.TotalDiscount = 0
	orderModel.PurchasePrice = &finalPrice
	now := time.Now()
	orderModel.Date = &now
	orderModel.ServiceUserID = &user.ServiceUser.ID
	if err := helpers.Convert(orderData, &orderModel); err != nil {
		return orderEntity, err
	}
	result, err := domains.RepoRegistry.ServiceOrderRepo.Create(c, orderModel)
	if err != nil {
		return orderEntity, err
	}
	if err := helpers.Convert(result, &orderEntity); err != nil {
		return orderEntity, err
	}
	return orderEntity, nil
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
