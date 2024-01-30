//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2024 Red Hat, Inc.
 *
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	configv1 "github.com/openshift/api/config/v1"
	apicorev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	corev1 "kubevirt.io/api/core/v1"
	v1alpha1 "kubevirt.io/applications-aware-quota/staging/src/kubevirt.io/applications-aware-quota-api/pkg/apis/core/v1alpha1"
	corev1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationAwareConfigurations) DeepCopyInto(out *ApplicationAwareConfigurations) {
	*out = *in
	if in.VmiCalcConfigName != nil {
		in, out := &in.VmiCalcConfigName, &out.VmiCalcConfigName
		*out = new(v1alpha1.VmiCalcConfigName)
		**out = **in
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationAwareConfigurations.
func (in *ApplicationAwareConfigurations) DeepCopy() *ApplicationAwareConfigurations {
	if in == nil {
		return nil
	}
	out := new(ApplicationAwareConfigurations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertRotateConfigCA) DeepCopyInto(out *CertRotateConfigCA) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.RenewBefore != nil {
		in, out := &in.RenewBefore, &out.RenewBefore
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertRotateConfigCA.
func (in *CertRotateConfigCA) DeepCopy() *CertRotateConfigCA {
	if in == nil {
		return nil
	}
	out := new(CertRotateConfigCA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertRotateConfigServer) DeepCopyInto(out *CertRotateConfigServer) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.RenewBefore != nil {
		in, out := &in.RenewBefore, &out.RenewBefore
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertRotateConfigServer.
func (in *CertRotateConfigServer) DeepCopy() *CertRotateConfigServer {
	if in == nil {
		return nil
	}
	out := new(CertRotateConfigServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataImportCronStatus) DeepCopyInto(out *DataImportCronStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataImportCronStatus.
func (in *DataImportCronStatus) DeepCopy() *DataImportCronStatus {
	if in == nil {
		return nil
	}
	out := new(DataImportCronStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataImportCronTemplate) DeepCopyInto(out *DataImportCronTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(corev1beta1.DataImportCronSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataImportCronTemplate.
func (in *DataImportCronTemplate) DeepCopy() *DataImportCronTemplate {
	if in == nil {
		return nil
	}
	out := new(DataImportCronTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataImportCronTemplateStatus) DeepCopyInto(out *DataImportCronTemplateStatus) {
	*out = *in
	in.DataImportCronTemplate.DeepCopyInto(&out.DataImportCronTemplate)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataImportCronTemplateStatus.
func (in *DataImportCronTemplateStatus) DeepCopy() *DataImportCronTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(DataImportCronTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConverged) DeepCopyInto(out *HyperConverged) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConverged.
func (in *HyperConverged) DeepCopy() *HyperConverged {
	if in == nil {
		return nil
	}
	out := new(HyperConverged)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HyperConverged) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedCertConfig) DeepCopyInto(out *HyperConvergedCertConfig) {
	*out = *in
	in.CA.DeepCopyInto(&out.CA)
	in.Server.DeepCopyInto(&out.Server)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedCertConfig.
func (in *HyperConvergedCertConfig) DeepCopy() *HyperConvergedCertConfig {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedCertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedConfig) DeepCopyInto(out *HyperConvergedConfig) {
	*out = *in
	if in.NodePlacement != nil {
		in, out := &in.NodePlacement, &out.NodePlacement
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedConfig.
func (in *HyperConvergedConfig) DeepCopy() *HyperConvergedConfig {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedFeatureGates) DeepCopyInto(out *HyperConvergedFeatureGates) {
	*out = *in
	if in.WithHostPassthroughCPU != nil {
		in, out := &in.WithHostPassthroughCPU, &out.WithHostPassthroughCPU
		*out = new(bool)
		**out = **in
	}
	if in.EnableCommonBootImageImport != nil {
		in, out := &in.EnableCommonBootImageImport, &out.EnableCommonBootImageImport
		*out = new(bool)
		**out = **in
	}
	if in.DeployTektonTaskResources != nil {
		in, out := &in.DeployTektonTaskResources, &out.DeployTektonTaskResources
		*out = new(bool)
		**out = **in
	}
	if in.DeployVMConsoleProxy != nil {
		in, out := &in.DeployVMConsoleProxy, &out.DeployVMConsoleProxy
		*out = new(bool)
		**out = **in
	}
	if in.DeployKubeSecondaryDNS != nil {
		in, out := &in.DeployKubeSecondaryDNS, &out.DeployKubeSecondaryDNS
		*out = new(bool)
		**out = **in
	}
	if in.NonRoot != nil {
		in, out := &in.NonRoot, &out.NonRoot
		*out = new(bool)
		**out = **in
	}
	if in.DisableMDevConfiguration != nil {
		in, out := &in.DisableMDevConfiguration, &out.DisableMDevConfiguration
		*out = new(bool)
		**out = **in
	}
	if in.PersistentReservation != nil {
		in, out := &in.PersistentReservation, &out.PersistentReservation
		*out = new(bool)
		**out = **in
	}
	if in.EnableManagedTenantQuota != nil {
		in, out := &in.EnableManagedTenantQuota, &out.EnableManagedTenantQuota
		*out = new(bool)
		**out = **in
	}
	if in.AutoResourceLimits != nil {
		in, out := &in.AutoResourceLimits, &out.AutoResourceLimits
		*out = new(bool)
		**out = **in
	}
	if in.AlignCPUs != nil {
		in, out := &in.AlignCPUs, &out.AlignCPUs
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedFeatureGates.
func (in *HyperConvergedFeatureGates) DeepCopy() *HyperConvergedFeatureGates {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedFeatureGates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedList) DeepCopyInto(out *HyperConvergedList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HyperConverged, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedList.
func (in *HyperConvergedList) DeepCopy() *HyperConvergedList {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HyperConvergedList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedObsoleteCPUs) DeepCopyInto(out *HyperConvergedObsoleteCPUs) {
	*out = *in
	if in.CPUModels != nil {
		in, out := &in.CPUModels, &out.CPUModels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedObsoleteCPUs.
func (in *HyperConvergedObsoleteCPUs) DeepCopy() *HyperConvergedObsoleteCPUs {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedObsoleteCPUs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedSpec) DeepCopyInto(out *HyperConvergedSpec) {
	*out = *in
	in.Infra.DeepCopyInto(&out.Infra)
	in.Workloads.DeepCopyInto(&out.Workloads)
	in.FeatureGates.DeepCopyInto(&out.FeatureGates)
	in.LiveMigrationConfig.DeepCopyInto(&out.LiveMigrationConfig)
	if in.PermittedHostDevices != nil {
		in, out := &in.PermittedHostDevices, &out.PermittedHostDevices
		*out = new(PermittedHostDevices)
		(*in).DeepCopyInto(*out)
	}
	if in.MediatedDevicesConfiguration != nil {
		in, out := &in.MediatedDevicesConfiguration, &out.MediatedDevicesConfiguration
		*out = new(MediatedDevicesConfiguration)
		(*in).DeepCopyInto(*out)
	}
	in.CertConfig.DeepCopyInto(&out.CertConfig)
	if in.ResourceRequirements != nil {
		in, out := &in.ResourceRequirements, &out.ResourceRequirements
		*out = new(OperandResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ScratchSpaceStorageClass != nil {
		in, out := &in.ScratchSpaceStorageClass, &out.ScratchSpaceStorageClass
		*out = new(string)
		**out = **in
	}
	if in.VddkInitImage != nil {
		in, out := &in.VddkInitImage, &out.VddkInitImage
		*out = new(string)
		**out = **in
	}
	if in.DefaultCPUModel != nil {
		in, out := &in.DefaultCPUModel, &out.DefaultCPUModel
		*out = new(string)
		**out = **in
	}
	if in.DefaultRuntimeClass != nil {
		in, out := &in.DefaultRuntimeClass, &out.DefaultRuntimeClass
		*out = new(string)
		**out = **in
	}
	if in.ObsoleteCPUs != nil {
		in, out := &in.ObsoleteCPUs, &out.ObsoleteCPUs
		*out = new(HyperConvergedObsoleteCPUs)
		(*in).DeepCopyInto(*out)
	}
	if in.CommonTemplatesNamespace != nil {
		in, out := &in.CommonTemplatesNamespace, &out.CommonTemplatesNamespace
		*out = new(string)
		**out = **in
	}
	if in.StorageImport != nil {
		in, out := &in.StorageImport, &out.StorageImport
		*out = new(StorageImportConfig)
		(*in).DeepCopyInto(*out)
	}
	in.WorkloadUpdateStrategy.DeepCopyInto(&out.WorkloadUpdateStrategy)
	if in.DataImportCronTemplates != nil {
		in, out := &in.DataImportCronTemplates, &out.DataImportCronTemplates
		*out = make([]DataImportCronTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FilesystemOverhead != nil {
		in, out := &in.FilesystemOverhead, &out.FilesystemOverhead
		*out = new(corev1beta1.FilesystemOverhead)
		(*in).DeepCopyInto(*out)
	}
	if in.LogVerbosityConfig != nil {
		in, out := &in.LogVerbosityConfig, &out.LogVerbosityConfig
		*out = new(LogVerbosityConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSSecurityProfile != nil {
		in, out := &in.TLSSecurityProfile, &out.TLSSecurityProfile
		*out = new(configv1.TLSSecurityProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.TektonPipelinesNamespace != nil {
		in, out := &in.TektonPipelinesNamespace, &out.TektonPipelinesNamespace
		*out = new(string)
		**out = **in
	}
	if in.TektonTasksNamespace != nil {
		in, out := &in.TektonTasksNamespace, &out.TektonTasksNamespace
		*out = new(string)
		**out = **in
	}
	if in.KubeSecondaryDNSNameServerIP != nil {
		in, out := &in.KubeSecondaryDNSNameServerIP, &out.KubeSecondaryDNSNameServerIP
		*out = new(string)
		**out = **in
	}
	if in.EvictionStrategy != nil {
		in, out := &in.EvictionStrategy, &out.EvictionStrategy
		*out = new(corev1.EvictionStrategy)
		**out = **in
	}
	if in.VMStateStorageClass != nil {
		in, out := &in.VMStateStorageClass, &out.VMStateStorageClass
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineOptions != nil {
		in, out := &in.VirtualMachineOptions, &out.VirtualMachineOptions
		*out = new(VirtualMachineOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.CommonBootImageNamespace != nil {
		in, out := &in.CommonBootImageNamespace, &out.CommonBootImageNamespace
		*out = new(string)
		**out = **in
	}
	if in.KSMConfiguration != nil {
		in, out := &in.KSMConfiguration, &out.KSMConfiguration
		*out = new(corev1.KSMConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkBinding != nil {
		in, out := &in.NetworkBinding, &out.NetworkBinding
		*out = make(map[string]corev1.InterfaceBindingPlugin, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ApplicationAwareConfig != nil {
		in, out := &in.ApplicationAwareConfig, &out.ApplicationAwareConfig
		*out = new(ApplicationAwareConfigurations)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedSpec.
func (in *HyperConvergedSpec) DeepCopy() *HyperConvergedSpec {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedStatus) DeepCopyInto(out *HyperConvergedStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]apicorev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]Version, len(*in))
		copy(*out, *in)
	}
	if in.DataImportCronTemplates != nil {
		in, out := &in.DataImportCronTemplates, &out.DataImportCronTemplates
		*out = make([]DataImportCronTemplateStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedStatus.
func (in *HyperConvergedStatus) DeepCopy() *HyperConvergedStatus {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperConvergedWorkloadUpdateStrategy) DeepCopyInto(out *HyperConvergedWorkloadUpdateStrategy) {
	*out = *in
	if in.WorkloadUpdateMethods != nil {
		in, out := &in.WorkloadUpdateMethods, &out.WorkloadUpdateMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BatchEvictionSize != nil {
		in, out := &in.BatchEvictionSize, &out.BatchEvictionSize
		*out = new(int)
		**out = **in
	}
	if in.BatchEvictionInterval != nil {
		in, out := &in.BatchEvictionInterval, &out.BatchEvictionInterval
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperConvergedWorkloadUpdateStrategy.
func (in *HyperConvergedWorkloadUpdateStrategy) DeepCopy() *HyperConvergedWorkloadUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(HyperConvergedWorkloadUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LiveMigrationConfigurations) DeepCopyInto(out *LiveMigrationConfigurations) {
	*out = *in
	if in.ParallelMigrationsPerCluster != nil {
		in, out := &in.ParallelMigrationsPerCluster, &out.ParallelMigrationsPerCluster
		*out = new(uint32)
		**out = **in
	}
	if in.ParallelOutboundMigrationsPerNode != nil {
		in, out := &in.ParallelOutboundMigrationsPerNode, &out.ParallelOutboundMigrationsPerNode
		*out = new(uint32)
		**out = **in
	}
	if in.BandwidthPerMigration != nil {
		in, out := &in.BandwidthPerMigration, &out.BandwidthPerMigration
		*out = new(string)
		**out = **in
	}
	if in.CompletionTimeoutPerGiB != nil {
		in, out := &in.CompletionTimeoutPerGiB, &out.CompletionTimeoutPerGiB
		*out = new(int64)
		**out = **in
	}
	if in.ProgressTimeout != nil {
		in, out := &in.ProgressTimeout, &out.ProgressTimeout
		*out = new(int64)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.AllowAutoConverge != nil {
		in, out := &in.AllowAutoConverge, &out.AllowAutoConverge
		*out = new(bool)
		**out = **in
	}
	if in.AllowPostCopy != nil {
		in, out := &in.AllowPostCopy, &out.AllowPostCopy
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LiveMigrationConfigurations.
func (in *LiveMigrationConfigurations) DeepCopy() *LiveMigrationConfigurations {
	if in == nil {
		return nil
	}
	out := new(LiveMigrationConfigurations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogVerbosityConfiguration) DeepCopyInto(out *LogVerbosityConfiguration) {
	*out = *in
	if in.Kubevirt != nil {
		in, out := &in.Kubevirt, &out.Kubevirt
		*out = new(corev1.LogVerbosity)
		(*in).DeepCopyInto(*out)
	}
	if in.CDI != nil {
		in, out := &in.CDI, &out.CDI
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogVerbosityConfiguration.
func (in *LogVerbosityConfiguration) DeepCopy() *LogVerbosityConfiguration {
	if in == nil {
		return nil
	}
	out := new(LogVerbosityConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MediatedDevicesConfiguration) DeepCopyInto(out *MediatedDevicesConfiguration) {
	*out = *in
	if in.MediatedDeviceTypes != nil {
		in, out := &in.MediatedDeviceTypes, &out.MediatedDeviceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MediatedDevicesTypes != nil {
		in, out := &in.MediatedDevicesTypes, &out.MediatedDevicesTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeMediatedDeviceTypes != nil {
		in, out := &in.NodeMediatedDeviceTypes, &out.NodeMediatedDeviceTypes
		*out = make([]NodeMediatedDeviceTypesConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MediatedDevicesConfiguration.
func (in *MediatedDevicesConfiguration) DeepCopy() *MediatedDevicesConfiguration {
	if in == nil {
		return nil
	}
	out := new(MediatedDevicesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MediatedHostDevice) DeepCopyInto(out *MediatedHostDevice) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MediatedHostDevice.
func (in *MediatedHostDevice) DeepCopy() *MediatedHostDevice {
	if in == nil {
		return nil
	}
	out := new(MediatedHostDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMediatedDeviceTypesConfig) DeepCopyInto(out *NodeMediatedDeviceTypesConfig) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MediatedDeviceTypes != nil {
		in, out := &in.MediatedDeviceTypes, &out.MediatedDeviceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MediatedDevicesTypes != nil {
		in, out := &in.MediatedDevicesTypes, &out.MediatedDevicesTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMediatedDeviceTypesConfig.
func (in *NodeMediatedDeviceTypesConfig) DeepCopy() *NodeMediatedDeviceTypesConfig {
	if in == nil {
		return nil
	}
	out := new(NodeMediatedDeviceTypesConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperandResourceRequirements) DeepCopyInto(out *OperandResourceRequirements) {
	*out = *in
	if in.StorageWorkloads != nil {
		in, out := &in.StorageWorkloads, &out.StorageWorkloads
		*out = new(apicorev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.VmiCPUAllocationRatio != nil {
		in, out := &in.VmiCPUAllocationRatio, &out.VmiCPUAllocationRatio
		*out = new(int)
		**out = **in
	}
	if in.AutoCPULimitNamespaceLabelSelector != nil {
		in, out := &in.AutoCPULimitNamespaceLabelSelector, &out.AutoCPULimitNamespaceLabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperandResourceRequirements.
func (in *OperandResourceRequirements) DeepCopy() *OperandResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(OperandResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PciHostDevice) DeepCopyInto(out *PciHostDevice) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PciHostDevice.
func (in *PciHostDevice) DeepCopy() *PciHostDevice {
	if in == nil {
		return nil
	}
	out := new(PciHostDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermittedHostDevices) DeepCopyInto(out *PermittedHostDevices) {
	*out = *in
	if in.PciHostDevices != nil {
		in, out := &in.PciHostDevices, &out.PciHostDevices
		*out = make([]PciHostDevice, len(*in))
		copy(*out, *in)
	}
	if in.MediatedDevices != nil {
		in, out := &in.MediatedDevices, &out.MediatedDevices
		*out = make([]MediatedHostDevice, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermittedHostDevices.
func (in *PermittedHostDevices) DeepCopy() *PermittedHostDevices {
	if in == nil {
		return nil
	}
	out := new(PermittedHostDevices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageImportConfig) DeepCopyInto(out *StorageImportConfig) {
	*out = *in
	if in.InsecureRegistries != nil {
		in, out := &in.InsecureRegistries, &out.InsecureRegistries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageImportConfig.
func (in *StorageImportConfig) DeepCopy() *StorageImportConfig {
	if in == nil {
		return nil
	}
	out := new(StorageImportConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Version) DeepCopyInto(out *Version) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Version.
func (in *Version) DeepCopy() *Version {
	if in == nil {
		return nil
	}
	out := new(Version)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineOptions) DeepCopyInto(out *VirtualMachineOptions) {
	*out = *in
	if in.DisableFreePageReporting != nil {
		in, out := &in.DisableFreePageReporting, &out.DisableFreePageReporting
		*out = new(bool)
		**out = **in
	}
	if in.DisableSerialConsoleLog != nil {
		in, out := &in.DisableSerialConsoleLog, &out.DisableSerialConsoleLog
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineOptions.
func (in *VirtualMachineOptions) DeepCopy() *VirtualMachineOptions {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineOptions)
	in.DeepCopyInto(out)
	return out
}
