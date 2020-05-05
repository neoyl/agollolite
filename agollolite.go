package agollolite

import (
	"github.com/shima-park/agollo"
)

type Client interface {
	GetString(namespace, key string) string
	GetUint16(namespace string, key string) (uint16, error)
}

type client struct {
	agolloClient agollo.Agollo
}

func New(metaServerUrl, appId, cluster string) (Client, error) {
	agolloClient, err := agollo.New(metaServerUrl, appId, agollo.AutoFetchOnCacheMiss(), agollo.Cluster(cluster))
	if nil != err {
		return nil, err
	}
	return client{agolloClient: agolloClient}, nil
}
