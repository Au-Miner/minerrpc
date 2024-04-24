package provider

import "reflect"

type ServiceProvider interface {
	AddServiceProvider(iClass interface{})
	GetFunc(funcName string) (reflect.Value, error)
}
