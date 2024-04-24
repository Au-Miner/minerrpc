package provider

import (
	"fmt"
	"reflect"
)

type ServiceProviderImpl struct {
	funcs map[string]reflect.Value
}

func NewServiceProvider() *ServiceProviderImpl {
	return &ServiceProviderImpl{funcs: make(map[string]reflect.Value)}
}

func (sp *ServiceProviderImpl) AddServiceProvider(iClass interface{}) {
	rType := reflect.TypeOf(iClass)
	rClass := reflect.ValueOf(iClass)
	for idx := 0; idx < rClass.NumMethod(); idx++ {
		rFuncType := rType.Method(idx)
		rFuncClass := rClass.Method(idx)
		if rFuncType.Type.Kind() == reflect.Func {
			if _, ok := sp.funcs[rFuncType.Name]; ok {
				continue
			}
			sp.funcs[rFuncType.Name] = rFuncClass
		}
	}
}

func (sp *ServiceProviderImpl) GetFunc(funcName string) (reflect.Value, error) {
	f, ok := sp.funcs[funcName]
	if !ok {
		return reflect.Value{}, fmt.Errorf("func not found")
	} else {
		return f, nil
	}
}
