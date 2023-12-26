package services

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

type orderService struct{}

func NewOrderService() domains.OrderService {
	return &orderService{}
}

func (orderService) MakeOrder(c context.Context, user domains.UserEntity, orderData domains.ServiceOrderForm) (domains.ServiceOrderEntity, error) {
	var orderEntity domains.ServiceOrderEntity
	var orderModel domains.ServiceOrderModel
	details := orderData.Details
	orderData.Details = nil
	if ok, err := helpers.Validate(orderData); !ok {
		return domains.ServiceOrderEntity{}, err
	}
	var orderTotalPrice uint64
	var validatedDetail []domains.ServiceOrderDetailModel
	for _, detail := range details {
		detailItems := detail.Items
		detail.Items = nil
		if ok, err := helpers.Validate(detail); !ok {
			return domains.ServiceOrderEntity{}, err
		}
		psId := detail.PartialServiceID
		ps, err := domains.RepoRegistry.ServiceRepo.Find(c, psId)
		if err != nil {
			conv := domains.ErrForeignKeyViolated
			return orderEntity, conv
		}
		var additionalPrice uint64
		detailData := domains.ServiceOrderDetailModel{
			PartialServiceID: &ps.ID,
			ServicePrice:     ps.BasePrice,
		}
		itemList := []domains.ServiceOrderDetailItemModel{}
		for _, item := range detailItems {
			if ok, err := helpers.Validate(item); !ok {
				return domains.ServiceOrderEntity{}, err
			}
			itemId := item.PartialServiceItemID
			si, err := domains.RepoRegistry.ServiceItemRepo.Find(c, itemId)
			if err != nil {
				conv := domains.ErrForeignKeyViolated
				return orderEntity, conv
			}
			if si.Service.ID != ps.ID {
				conv := domains.ErrUnmatchedData
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
			itemList = append(itemList, detailDataItem)
		}
		detailData.PartialServiceItems = itemList
		detailTotalPrice := (additionalPrice + *detailData.ServicePrice) * uint64(orderData.ContractDuration)
		detailData.AdditionalPrice = &additionalPrice
		detailData.TotalPrice = &detailTotalPrice
		orderTotalPrice += detailTotalPrice
		validatedDetail = append(validatedDetail, detailData)
	}
	status := domains.WaitingMOU
	totalItem := uint(len(details))
	finalPrice := orderTotalPrice * uint64(orderData.ContractDuration)
	orderModel.TotalPrice = &finalPrice
	orderModel.TotalDiscount = 0
	orderModel.PurchasePrice = &finalPrice
	orderModel.TotalItem = &totalItem
	orderModel.Status = &status
	orderModel.ServicePackageID = nil
	now := time.Now()
	orderModel.Date = &now
	orderModel.ServiceUserID = &user.ServiceUser.ID
	if err := helpers.Convert(orderData, &orderModel); err != nil {
		return orderEntity, err
	}
	orderModel.ServiceOrderDetails = validatedDetail
	result, err := domains.RepoRegistry.ServiceOrderRepo.Create(c, orderModel)
	if err != nil {
		return orderEntity, err
	}
	if err := helpers.Convert(result, &orderEntity); err != nil {
		return orderEntity, err
	}
	return orderEntity, nil
}
func (orderService) CancelOrder(c context.Context, user domains.UserEntity, orderId uint) (bool, error) {
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
	user domains.UserEntity,
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
func (orderService) DetailOrder(c context.Context, user domains.UserEntity, id uint) (domains.ServiceOrderEntity, error) {
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
func (orderService) UpdateOrderStatus(c context.Context, user domains.UserEntity, id uint, data domains.ServiceOrderUpdateStatusForm) (int, domains.ServiceOrderEntity, error) {
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

func (orderService) UploadMOU(c context.Context, user domains.UserEntity, orderId uint, fileObj domains.EntityFileMap) (bool, error) {
	var dataModel domains.ServiceOrderModel
	var dataEntity domains.ServiceOrderEntity
	fun := func(id uint) (int, domains.Model, error) {
		order, err := domains.RepoRegistry.ServiceOrderRepo.Find(c, orderId)
		if err != nil {
			return 0, dataModel, err
		}
		status := order.Status
		acceptedStatus := []string{
			string(domains.WaitingForConfirmationOrderStatus),
			string(domains.WaitingMOU),
		}
		if !slices.Contains(acceptedStatus, string(*status)) {
			return 0, dataModel, domains.ErrUnprocessAbleEntity
		}
		file := fileObj.File
		if file == nil {
			err := domains.ErrValidation
			err.ValidationErrors = govalidator.Errors{
				govalidator.Error{
					Name:                     "mou",
					Validator:                "required",
					CustomErrorMessageExists: true,
					Err:                      errors.New("mou file is required"),
				},
			}
			return 0, dataModel, err
		}
		mapped := map[string]domains.FileWrapper{
			"mou": {
				File:   file,
				Dest:   configs.FILE_DESTS["order/mou"],
				Field:  "mou",
				Config: configs.PDF_FILE_CONFIG,
			},
		}
		saveResult, _, err := domains.ServiceRegistry.FileServ.BatchStore(mapped)
		if err != nil {
			return 0, dataModel, err
		}
		old := dataModel.MOU
		go domains.ServiceRegistry.FileServ.Destroy(old)
		dataModel = order
		dataModel.ID = 0
		dataModel.MOU = saveResult["mou"]
		toStatus := domains.WaitingForConfirmationOrderStatus
		dataModel.Status = &toStatus
		aff, updated, err := domains.RepoRegistry.ServiceOrderRepo.Update(c, orderId, dataModel)
		return int(aff), updated, err
	}
	_, _, err := basicUpdateService(orderId, domains.ServiceOrderEntity{}, &dataModel, &dataEntity, fun)
	if err != nil {
		return false, err
	}
	return true, nil
}
