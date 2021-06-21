package agollolite

import (
	"strconv"

	"github.com/shima-park/agollo"
)

func (c client) GetString(namespace, key string) string {
	return c.agolloClient.Get(key, agollo.WithNamespace(namespace))
}

func (c client) GetUint8(namespace string, key string) (uint8, error) {
	strValue := c.GetString(namespace, key)
	value, err := strconv.ParseUint(strValue, 10, 8)
	if nil != err {
		return 0, err
	} else {
		return uint8(value), nil
	}
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

func (c client) GetInt(namespace string, key string) (int, error) {
	strValue := c.GetString(namespace, key)
	value, err := strconv.Atoi(strValue)
	if nil != err {
		return 0, err
	} else {
		return value, nil
	}
}
