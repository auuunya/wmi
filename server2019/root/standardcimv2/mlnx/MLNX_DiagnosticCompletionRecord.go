// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_DiagnosticCompletionRecord struct
type MLNX_DiagnosticCompletionRecord struct {
	*CIM_DiagnosticCompletionRecord
}

func NewMLNX_DiagnosticCompletionRecordEx1(instance *cim.WmiInstance) (newInstance *MLNX_DiagnosticCompletionRecord, err error) {
	tmp, err := NewCIM_DiagnosticCompletionRecordEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticCompletionRecord{
		CIM_DiagnosticCompletionRecord: tmp,
	}
	return
}

func NewMLNX_DiagnosticCompletionRecordEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DiagnosticCompletionRecord, err error) {
	tmp, err := NewCIM_DiagnosticCompletionRecordEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DiagnosticCompletionRecord{
		CIM_DiagnosticCompletionRecord: tmp,
	}
	return
}
