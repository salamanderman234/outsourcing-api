package helpers

import (
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

//SAMPLE REQUEST START HERE

// 1. Initiate Snap client

type midtransHelper struct{}

func (midtransHelper) CreatePayment(
	orderID string,
	totalAmount int64,
	items []models.TransactionDetail,
	userClient models.ServiceUser,
) (*snap.Response, *midtrans.Error) {
	client := providers.MidtransClient
	// client data
	empty := ""
	if userClient.UserID == nil {
		zero := uint(0)
		userClient.UserID = &zero
	}
	userID := strconv.Itoa(int(*userClient.UserID))
	if userClient.Fullname == nil {
		userClient.Fullname = &empty
	}
	fullname := *userClient.Fullname
	if userClient.User.Email == nil {
		userClient.User.Email = &empty
	}
	email := *userClient.User.Email
	if userClient.User.Email == nil {
		userClient.User.Email = &empty
	}
	phone := *userClient.Phone
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: totalAmount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},

		CustomerDetail: &midtrans.CustomerDetails{
			FName: fullname,
			Email: email,
			Phone: phone,
		},
		UserId: userID,
		Expiry: &snap.ExpiryDetails{
			Duration: 30,
			Unit:     "minutes",
		},
	}
	resp, err := client.CreateTransaction(req)
	return resp, err
}

var Midtrans = midtransHelper{}
