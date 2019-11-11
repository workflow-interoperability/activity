package lib

import "github.com/zeebe-io/zeebe/clients/go/zbc"

func ConnectZeebe() (*zbc.ZBClient, error) {
	const BrokerAddr = "0.0.0.0:26500"
	client, err := zbc.NewZBClientWithConfig(&zbc.ZBClientConfig{
		GatewayAddress:         BrokerAddr,
		UsePlaintextConnection: true,
	})
	return &client, err
}
