// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ParentChildSettingData struct
type Msvm_ParentChildSettingData struct {
	*CIM_Dependency
}

func NewMsvm_ParentChildSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_ParentChildSettingData, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ParentChildSettingData{
		CIM_Dependency: tmp,
	}
	return
}

func NewMsvm_ParentChildSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ParentChildSettingData, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ParentChildSettingData{
		CIM_Dependency: tmp,
	}
	return
}
