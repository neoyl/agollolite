package agollolite

import "github.com/shima-park/agollo"

type Client interface {
	GetString(namespace, key string) string
}

type client struct {
	agolloClient agollo.Agollo
}

func (c client) GetString(namespace, key string) string {
	return c.agolloClient.Get(key, agollo.WithNamespace(namespace))
}

func New(metaServerUrl, appId, cluster string) (Client, error) {
	agolloClient, err := agollo.New(metaServerUrl, appId, agollo.AutoFetchOnCacheMiss(), agollo.Cluster(cluster))
	if nil != err {
		return nil, err
	}
	return client{agolloClient: agolloClient}, nil
}
