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
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_ComputerSystem struct
type Msvm_ComputerSystem struct {
	*CIM_ComputerSystem

	//
	EnhancedSessionModeState uint16

	//
	FailedOverReplicationType uint16

	//
	HwThreadsPerCoreRealized uint32

	//
	LastApplicationConsistentReplicationTime string

	//
	LastReplicationTime string

	//
	LastReplicationType uint16

	//
	LastSuccessfulBackupTime string

	//
	NumberOfNumaNodes uint16

	//
	OnTimeInMilliseconds uint64

	//
	ProcessID uint32

	//
	ReplicationHealth uint16

	//
	ReplicationMode uint16

	//
	ReplicationState uint16

	//
	TimeOfLastConfigurationChange string
}

func NewMsvm_ComputerSystemEx1(instance *cim.WmiInstance) (newInstance *Msvm_ComputerSystem, err error) {
	tmp, err := NewCIM_ComputerSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ComputerSystem{
		CIM_ComputerSystem: tmp,
	}
	return
}

func NewMsvm_ComputerSystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ComputerSystem, err error) {
	tmp, err := NewCIM_ComputerSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ComputerSystem{
		CIM_ComputerSystem: tmp,
	}
	return
}

// SetEnhancedSessionModeState sets the value of EnhancedSessionModeState for the instance
func (instance *Msvm_ComputerSystem) SetPropertyEnhancedSessionModeState(value uint16) (err error) {
	return instance.SetProperty("EnhancedSessionModeState", (value))
}

// GetEnhancedSessionModeState gets the value of EnhancedSessionModeState for the instance
func (instance *Msvm_ComputerSystem) GetPropertyEnhancedSessionModeState() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnhancedSessionModeState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetFailedOverReplicationType sets the value of FailedOverReplicationType for the instance
func (instance *Msvm_ComputerSystem) SetPropertyFailedOverReplicationType(value uint16) (err error) {
	return instance.SetProperty("FailedOverReplicationType", (value))
}

// GetFailedOverReplicationType gets the value of FailedOverReplicationType for the instance
func (instance *Msvm_ComputerSystem) GetPropertyFailedOverReplicationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("FailedOverReplicationType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetHwThreadsPerCoreRealized sets the value of HwThreadsPerCoreRealized for the instance
func (instance *Msvm_ComputerSystem) SetPropertyHwThreadsPerCoreRealized(value uint32) (err error) {
	return instance.SetProperty("HwThreadsPerCoreRealized", (value))
}

// GetHwThreadsPerCoreRealized gets the value of HwThreadsPerCoreRealized for the instance
func (instance *Msvm_ComputerSystem) GetPropertyHwThreadsPerCoreRealized() (value uint32, err error) {
	retValue, err := instance.GetProperty("HwThreadsPerCoreRealized")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetLastApplicationConsistentReplicationTime sets the value of LastApplicationConsistentReplicationTime for the instance
func (instance *Msvm_ComputerSystem) SetPropertyLastApplicationConsistentReplicationTime(value string) (err error) {
	return instance.SetProperty("LastApplicationConsistentReplicationTime", (value))
}

// GetLastApplicationConsistentReplicationTime gets the value of LastApplicationConsistentReplicationTime for the instance
func (instance *Msvm_ComputerSystem) GetPropertyLastApplicationConsistentReplicationTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastApplicationConsistentReplicationTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLastReplicationTime sets the value of LastReplicationTime for the instance
func (instance *Msvm_ComputerSystem) SetPropertyLastReplicationTime(value string) (err error) {
	return instance.SetProperty("LastReplicationTime", (value))
}

// GetLastReplicationTime gets the value of LastReplicationTime for the instance
func (instance *Msvm_ComputerSystem) GetPropertyLastReplicationTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastReplicationTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLastReplicationType sets the value of LastReplicationType for the instance
func (instance *Msvm_ComputerSystem) SetPropertyLastReplicationType(value uint16) (err error) {
	return instance.SetProperty("LastReplicationType", (value))
}

// GetLastReplicationType gets the value of LastReplicationType for the instance
func (instance *Msvm_ComputerSystem) GetPropertyLastReplicationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("LastReplicationType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetLastSuccessfulBackupTime sets the value of LastSuccessfulBackupTime for the instance
func (instance *Msvm_ComputerSystem) SetPropertyLastSuccessfulBackupTime(value string) (err error) {
	return instance.SetProperty("LastSuccessfulBackupTime", (value))
}

// GetLastSuccessfulBackupTime gets the value of LastSuccessfulBackupTime for the instance
func (instance *Msvm_ComputerSystem) GetPropertyLastSuccessfulBackupTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastSuccessfulBackupTime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetNumberOfNumaNodes sets the value of NumberOfNumaNodes for the instance
func (instance *Msvm_ComputerSystem) SetPropertyNumberOfNumaNodes(value uint16) (err error) {
	return instance.SetProperty("NumberOfNumaNodes", (value))
}

// GetNumberOfNumaNodes gets the value of NumberOfNumaNodes for the instance
func (instance *Msvm_ComputerSystem) GetPropertyNumberOfNumaNodes() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfNumaNodes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetOnTimeInMilliseconds sets the value of OnTimeInMilliseconds for the instance
func (instance *Msvm_ComputerSystem) SetPropertyOnTimeInMilliseconds(value uint64) (err error) {
	return instance.SetProperty("OnTimeInMilliseconds", (value))
}

// GetOnTimeInMilliseconds gets the value of OnTimeInMilliseconds for the instance
func (instance *Msvm_ComputerSystem) GetPropertyOnTimeInMilliseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("OnTimeInMilliseconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetProcessID sets the value of ProcessID for the instance
func (instance *Msvm_ComputerSystem) SetPropertyProcessID(value uint32) (err error) {
	return instance.SetProperty("ProcessID", (value))
}

// GetProcessID gets the value of ProcessID for the instance
func (instance *Msvm_ComputerSystem) GetPropertyProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetReplicationHealth sets the value of ReplicationHealth for the instance
func (instance *Msvm_ComputerSystem) SetPropertyReplicationHealth(value uint16) (err error) {
	return instance.SetProperty("ReplicationHealth", (value))
}

// GetReplicationHealth gets the value of ReplicationHealth for the instance
func (instance *Msvm_ComputerSystem) GetPropertyReplicationHealth() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationHealth")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetReplicationMode sets the value of ReplicationMode for the instance
func (instance *Msvm_ComputerSystem) SetPropertyReplicationMode(value uint16) (err error) {
	return instance.SetProperty("ReplicationMode", (value))
}

// GetReplicationMode gets the value of ReplicationMode for the instance
func (instance *Msvm_ComputerSystem) GetPropertyReplicationMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetReplicationState sets the value of ReplicationState for the instance
func (instance *Msvm_ComputerSystem) SetPropertyReplicationState(value uint16) (err error) {
	return instance.SetProperty("ReplicationState", (value))
}

// GetReplicationState gets the value of ReplicationState for the instance
func (instance *Msvm_ComputerSystem) GetPropertyReplicationState() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetTimeOfLastConfigurationChange sets the value of TimeOfLastConfigurationChange for the instance
func (instance *Msvm_ComputerSystem) SetPropertyTimeOfLastConfigurationChange(value string) (err error) {
	return instance.SetProperty("TimeOfLastConfigurationChange", (value))
}

// GetTimeOfLastConfigurationChange gets the value of TimeOfLastConfigurationChange for the instance
func (instance *Msvm_ComputerSystem) GetPropertyTimeOfLastConfigurationChange() (value string, err error) {
	retValue, err := instance.GetProperty("TimeOfLastConfigurationChange")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

//

// <param name="RequestedState" type="uint16 "></param>
// <param name="TimeoutPeriod" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ComputerSystem) RequestReplicationStateChange( /* IN */ RequestedState uint16,
	/* OUT */ Job CIM_ConcreteJob,
	/* OPTIONAL IN */ TimeoutPeriod string,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RequestReplicationStateChange", Action, PercentComplete, Timeout, RequestedState, TimeoutPeriod)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Job" type="CIM_ConcreteJob ">May contain a reference to the ConcreteJob created to track the status of the interrupt injection.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ComputerSystem) InjectNonMaskableInterrupt( /* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("InjectNonMaskableInterrupt", Action, PercentComplete, Timeout)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReplicationRelationship" type="string "></param>
// <param name="RequestedState" type="uint16 "></param>
// <param name="TimeoutPeriod" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ComputerSystem) RequestReplicationStateChangeEx( /* IN */ ReplicationRelationship string,
	/* IN */ RequestedState uint16,
	/* OUT */ Job CIM_ConcreteJob,
	/* OPTIONAL IN */ TimeoutPeriod string,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RequestReplicationStateChangeEx", Action, PercentComplete, Timeout, ReplicationRelationship, RequestedState, TimeoutPeriod)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_ComputerSystem) GetRelatedRegisteredProfile() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_RegisteredProfile")
}

func (instance *Msvm_ComputerSystem) GetRelatedResourcePool() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_ResourcePool")
}

func (instance *Msvm_ComputerSystem) GetRelatedSynth3dVideoPool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_Synth3dVideoPool")
}

func (instance *Msvm_ComputerSystem) GetRelatedProcessorPool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ProcessorPool")
}

func (instance *Msvm_ComputerSystem) GetRelatedTerminalService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_TerminalService")
}

func (instance *Msvm_ComputerSystem) GetRelatedVirtualEthernetSwitchManagementService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitchManagementService")
}

func (instance *Msvm_ComputerSystem) GetRelatedImageManagementService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ImageManagementService")
}

func (instance *Msvm_ComputerSystem) GetRelatedReplicationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ReplicationService")
}

func (instance *Msvm_ComputerSystem) GetRelatedMetricService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_MetricService")
}

func (instance *Msvm_ComputerSystem) GetRelatedResourcePoolConfigurationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePoolConfigurationService")
}

func (instance *Msvm_ComputerSystem) GetRelatedSynthetic3DService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_Synthetic3DService")
}

func (instance *Msvm_ComputerSystem) GetRelatedAssignableDeviceService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AssignableDeviceService")
}

func (instance *Msvm_ComputerSystem) GetRelatedVirtualSystemManagementService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemManagementService")
}

func (instance *Msvm_ComputerSystem) GetRelatedVirtualSystemSnapshotService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemSnapshotService")
}

func (instance *Msvm_ComputerSystem) GetRelatedVirtualSystemMigrationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemMigrationService")
}

func (instance *Msvm_ComputerSystem) GetRelatedSecurityService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SecurityService")
}

func (instance *Msvm_ComputerSystem) GetRelatedCollectionManagementService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_CollectionManagementService")
}

func (instance *Msvm_ComputerSystem) GetRelatedCollectionSnapshotService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_CollectionSnapshotService")
}

func (instance *Msvm_ComputerSystem) GetRelatedVirtualSystemReferencePointService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemReferencePointService")
}

func (instance *Msvm_ComputerSystem) GetRelatedCollectionReferencePointService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_CollectionReferencePointService")
}

func (instance *Msvm_ComputerSystem) GetRelatedNumaNode() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_NumaNode")
}

func (instance *Msvm_ComputerSystem) GetRelatedInstalledEthernetSwitchExtension() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_InstalledEthernetSwitchExtension")
}

func (instance *Msvm_ComputerSystem) GetRelatedMemory() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_Memory")
}

func (instance *Msvm_ComputerSystem) GetRelatedProcessor() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_Processor")
}

func (instance *Msvm_ComputerSystem) GetRelatedInternalEthernetPort() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_InternalEthernetPort")
}

func (instance *Msvm_ComputerSystem) GetRelatedExternalEthernetPort() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_ExternalEthernetPort")
}

func (instance *Msvm_ComputerSystem) GetRelatedComputerSystem() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_ComputerSystem")
}
