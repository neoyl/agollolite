package agollolite

import (
	"io/ioutil"

	"github.com/shima-park/agollo"
	"gopkg.in/yaml.v3"
)

type Client interface {
	GetString(namespace, key string) string
	GetUint16(namespace string, key string) (uint16, error)
	GetInt(namespace string, key string) (int, error)
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

func NewWithConfigFile(configFile string) (Client, error) {
	f, err := ioutil.ReadFile(configFile)
	if nil != err {
		return nil, err
	}
	config := config{}
	err = yaml.Unmarshal(f, &config)
	if nil != err {
		return nil, err
	}
	return New(config.MetaServerUrl, config.AppId, config.Cluster)
}
