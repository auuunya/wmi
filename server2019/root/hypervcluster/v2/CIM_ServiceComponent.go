// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ServiceComponent struct
type CIM_ServiceComponent struct {
	*CIM_Component
}

func NewCIM_ServiceComponentEx1(instance *cim.WmiInstance) (newInstance *CIM_ServiceComponent, err error) {
	tmp, err := NewCIM_ComponentEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ServiceComponent{
		CIM_Component: tmp,
	}
	return
}

func NewCIM_ServiceComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ServiceComponent, err error) {
	tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ServiceComponent{
		CIM_Component: tmp,
	}
	return
}
