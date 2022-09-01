// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"fmt"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"

	"github.com/microsoft/wmi/server2019/root/cimv2"
)

type TotalProcessor struct {
	Cores             uint32
	LogicalProcessors uint32
}

type ProcessorInfo struct {
	Manufacturer      string
	Virtualization    bool
	ProcessorSpeed    int32
	CPUType           uint32
	Architecture      uint32
	Hypervisorpresent bool
}

type Processor struct {
	*cimv2.Win32_Processor
}

// NewPhysicalMemory
func NewProcessor(instance *wmi.WmiInstance) (*Processor, error) {
	wmivm, err := cimv2.NewWin32_ProcessorEx1(instance)
	if err != nil {
		return nil, err
	}
	return &Processor{wmivm}, nil
}

// GetTotalProcessor
func GetTotalProcessor(whost *host.WmiHost) (proc *TotalProcessor, err error) {
	query := query.NewWmiQuery("Win32_Processor")
	processors, err := instance.GetWmiInstancesFromHost(whost, string(constant.CimV2), query)
	if err != nil {
		return
	}
	defer processors.Close()

	totalCores := uint32(0)
	totalLogicalProcessors := uint32(0)

	for _, tmp := range processors {
		procInstance, err1 := cimv2.NewWin32_ProcessorEx1(tmp)
		if err1 != nil {
			err = err1
			return
		}

		cores, err1 := procInstance.GetProperty("NumberOfCores")
		if err1 != nil {
			err = err1
			return
		}
		totalCores = totalCores + uint32(cores.(int32))
		lp, err1 := procInstance.GetProperty("NumberOfLogicalProcessors")
		if err1 != nil {
			err = err1
			return
		}
		totalLogicalProcessors = totalLogicalProcessors + uint32(lp.(int32))
	}

	return &TotalProcessor{
		Cores:             totalCores,
		LogicalProcessors: totalLogicalProcessors,
	}, nil
}

// GetProcessorInfo
func GetProcessorInfo(whost *host.WmiHost) (proc *ProcessorInfo, err error) {
	queryProcessor := query.NewWmiQuery("Win32_Processor")
	queryComputerSystem := query.NewWmiQuery("Win32_ComputerSystem")

	procInfo, err := instance.GetWmiInstanceEx(whost, string(constant.CimV2), queryProcessor)
	if err != nil {
		return
	}
	defer procInfo.Close()

	winComputerSystemInfo, err := instance.GetWmiInstanceEx(whost, string(constant.CimV2), queryComputerSystem)
	if err != nil {
		return
	}
	defer winComputerSystemInfo.Close()

	procInstance, err := cimv2.NewWin32_ProcessorEx1(procInfo)
	if err != nil {
		return
	}

	manuf, err := procInstance.GetProperty("Manufacturer")
	if err != nil {
		return
	}

	isVirtualizationEnabled, err := procInstance.GetProperty("VirtualizationFirmwareEnabled")
	if err != nil {
		return
	}

	currClockSpeed, err := procInstance.GetProperty("CurrentClockSpeed")
	if err != nil {
		return
	}

	cpuType, err := procInstance.GetProperty("ProcessorType")
	if err != nil {
		return
	}

	architecture, err := procInstance.GetProperty("Architecture")
	if err != nil {
		return
	}

	winCompSystemInstance, err := cimv2.NewWin32_ComputerSystemEx1(winComputerSystemInfo)
	if err != nil {
		return
	}

	hypervisorPresent, err := winCompSystemInstance.GetProperty("HypervisorPresent")
	if err != nil {
		return
	}

	fmt.Printf("Manufacturer value in GetProcessorInfo is: [%v]\n ", manuf)
	fmt.Printf("Virtualization value in GetProcessorInfo is: [%v]\n ", isVirtualizationEnabled)
	fmt.Printf("HyperVisorPresent value in GetProcessorInfo is: [%v]\n ", hypervisorPresent)
	fmt.Printf("CurrentClock value in GetProcessorInfo is: [%v]\n ", currClockSpeed)
	fmt.Printf("ProcessorType value in GetProcessorInfo is: [%v]\n ", cpuType)
	fmt.Printf("Architecture value in GetProcessorInfo is: [%v]\n ", architecture)

	return &ProcessorInfo{
		Manufacturer:      manuf.(string),
		Virtualization:    isVirtualizationEnabled.(bool),
		Hypervisorpresent: hypervisorPresent.(bool),
		ProcessorSpeed:    currClockSpeed.(int32),
		CPUType:           uint32(cpuType.(int32)),
		Architecture:      uint32(architecture.(int32)),
	}, nil
}
