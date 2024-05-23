package providers

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/salamanderman234/outsourcing-api/configs"
)

var MidtransClient snap.Client

func SetMidtransClient() {
	serverKey := configs.MidtransConfig.ServerKey
	environment := configs.AppConfig.Env
	midtransEnv := midtrans.Sandbox
	if environment == "production" {
		midtransEnv = midtrans.Production
	}
	MidtransClient.New(serverKey, midtransEnv)
}
