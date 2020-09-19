// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PnPDevicePropertyReal32Array struct
type Win32_PnPDevicePropertyReal32Array struct {
	*Win32_PnPDeviceProperty

	//
	Data []float32
}

func NewWin32_PnPDevicePropertyReal32ArrayEx1(instance *cim.WmiInstance) (newInstance *Win32_PnPDevicePropertyReal32Array, err error) {
	tmp, err := NewWin32_PnPDevicePropertyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PnPDevicePropertyReal32Array{
		Win32_PnPDeviceProperty: tmp,
	}
	return
}

func NewWin32_PnPDevicePropertyReal32ArrayEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PnPDevicePropertyReal32Array, err error) {
	tmp, err := NewWin32_PnPDevicePropertyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PnPDevicePropertyReal32Array{
		Win32_PnPDeviceProperty: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyReal32Array) SetPropertyData(value []float32) (err error) {
	return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyReal32Array) GetPropertyData() (value []float32, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(float32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, float32(valuetmp))
	}

	return
}
