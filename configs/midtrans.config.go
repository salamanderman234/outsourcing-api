package configs

import (
	"github.com/spf13/viper"
)

type midConfigs struct {
	MerchantID string
	ClientKey  string
	ServerKey  string
}

func (m *midConfigs) setMidtransConfig() {
	m.MerchantID = viper.GetString("MIDTRANS_MERCHANT_ID")
	m.ClientKey = viper.GetString("MIDTRANS_CLIENT_KEY")
	m.ServerKey = viper.GetString("MIDTRANS_SERVER_KEY")
}

var MidtransConfig = midConfigs{}
