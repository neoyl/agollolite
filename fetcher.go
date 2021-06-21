package agollolite

import (
	"github.com/shima-park/agollo"
	"strconv"
)

func (c client) GetString(namespace, key string) string {
	return c.agolloClient.Get(key, agollo.WithNamespace(namespace))
}

func (c client) GetUint16(namespace string, key string) (uint16, error) {
	strValue := c.GetString(namespace, key)
	value, err := strconv.ParseUint(strValue, 10, 16)
	if nil != err {
		return 0, err
	} else {
		return uint16(value), nil
	}
}

func (c client) GetUint64(namespace string, key string) (uint64, error) {
	strValue := c.GetString(namespace, key)
	value, err := strconv.ParseUint(strValue, 10, 64)
	if nil != err {
		return 0, err
	} else {
		return value, nil
	}
}
