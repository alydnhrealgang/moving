package api

import (
	"github.com/alydnhrealgang/moving/cli/api/moving_clients"
	"net/url"
)

func CreateApiClient(rawUrl string) (*moving_clients.Moving, error) {
	movingUrl, err := url.Parse(rawUrl)
	if nil != err {
		return nil, err
	}
	return moving_clients.NewHTTPClientWithConfig(nil, &moving_clients.TransportConfig{
		Host:     movingUrl.Host,
		BasePath: movingUrl.Path,
		Schemes:  []string{movingUrl.Scheme},
	}), nil
}
