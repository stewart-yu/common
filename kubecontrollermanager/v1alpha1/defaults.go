/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	utilpointer "k8s.io/kubernetes/pkg/util/pointer"
)

func addDefaultingFuncs(scheme *kruntime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_KubeControllerManagerConfiguration(obj *KubeControllerManagerConfiguration) {
	zero := metav1.Duration{}
	if len(obj.Controllers) == 0 {
		obj.Controllers = []string{"*"}
	}
	if obj.AttachDetachController.ReconcilerSyncLoopPeriod == zero {
		obj.AttachDetachController.ReconcilerSyncLoopPeriod = metav1.Duration{Duration: 60 * time.Second}
	}
	if obj.DaemonSetController.ConcurrentDaemonSetSyncs == 0 {
		obj.DaemonSetController.ConcurrentDaemonSetSyncs = 2
	}
	if obj.DeprecatedController.RegisterRetryCount == 0 {
		obj.DeprecatedController.RegisterRetryCount = 10
	}
	if obj.EndPointController.ConcurrentEndpointSyncs == 0 {
		obj.EndPointController.ConcurrentEndpointSyncs = 5
	}
	if obj.JobController.ConcurrentJobSyncs == 0 {
		obj.JobController.ConcurrentJobSyncs = 5
	}
	if obj.NodeIpamController.NodeCIDRMaskSize == 0 {
		obj.NodeIpamController.NodeCIDRMaskSize = 24
	}
	if obj.PersistentVolumeBinderController.PVClaimBinderSyncPeriod == zero {
		obj.PersistentVolumeBinderController.PVClaimBinderSyncPeriod = metav1.Duration{Duration: 15 * time.Second}
	}
	if obj.PodGCController.TerminatedPodGCThreshold == 0 {
		obj.PodGCController.TerminatedPodGCThreshold = 12500
	}
	if obj.ReplicationController.ConcurrentRCSyncs == 0 {
		obj.ReplicationController.ConcurrentRCSyncs = 5
	}
	if obj.ReplicaSetController.ConcurrentRSSyncs == 0 {
		obj.ReplicaSetController.ConcurrentRSSyncs = 5
	}
	if obj.SAController.ConcurrentSATokenSyncs == 0 {
		obj.SAController.ConcurrentSATokenSyncs = 5
	}
}

func SetDefaults_CSRSigningControllerConfiguration(obj *CSRSigningControllerConfiguration) {
	zero := metav1.Duration{}
	if obj.ClusterSigningCertFile == "" {
		obj.ClusterSigningCertFile = "/etc/kubernetes/ca/ca.pem"
	}
	if obj.ClusterSigningKeyFile == "" {
		obj.ClusterSigningKeyFile = "/etc/kubernetes/ca/ca.key"
	}
	if obj.ClusterSigningDuration == zero {
		obj.ClusterSigningDuration = metav1.Duration{Duration: 365 * 24 * time.Hour}
	}
}

func SetDefaults_DeploymentControllerConfiguration(obj *DeploymentControllerConfiguration) {
	zero := metav1.Duration{}
	if obj.ConcurrentDeploymentSyncs == 0 {
		obj.ConcurrentDeploymentSyncs = 5
	}
	if obj.DeploymentControllerSyncPeriod == zero {
		obj.DeploymentControllerSyncPeriod = metav1.Duration{Duration: 30 * time.Second}
	}
}

func SetDefaults_GarbageCollectorControllerConfiguration(obj *GarbageCollectorControllerConfiguration) {
	if obj.EnableGarbageCollector == nil {
		obj.EnableGarbageCollector = utilpointer.BoolPtr(true)
	}
	if obj.ConcurrentGCSyncs == 0 {
		obj.ConcurrentGCSyncs = 20
	}
}

func SetDefaults_HPAControllerConfiguration(obj *HPAControllerConfiguration) {
	zero := metav1.Duration{}
	if obj.HorizontalPodAutoscalerSyncPeriod == zero {
		obj.HorizontalPodAutoscalerSyncPeriod = metav1.Duration{Duration: 30 * time.Second}
	}
	if obj.HorizontalPodAutoscalerUpscaleForbiddenWindow == zero {
		obj.HorizontalPodAutoscalerUpscaleForbiddenWindow = metav1.Duration{Duration: 3 * time.Minute}
	}
	if obj.HorizontalPodAutoscalerDownscaleForbiddenWindow == zero {
		obj.HorizontalPodAutoscalerDownscaleForbiddenWindow = metav1.Duration{Duration: 5 * time.Minute}
	}
	if obj.HorizontalPodAutoscalerTolerance == 0 {
		obj.HorizontalPodAutoscalerTolerance = 0.1
	}
	if obj.HorizontalPodAutoscalerUseRESTClients == nil {
		obj.HorizontalPodAutoscalerUseRESTClients = utilpointer.BoolPtr(true)
	}
}

func SetDefaults_NamespaceControllerConfiguration(obj *NamespaceControllerConfiguration) {
	zero := metav1.Duration{}
	if obj.ConcurrentNamespaceSyncs == 0 {
		obj.ConcurrentNamespaceSyncs = 10
	}
	if obj.NamespaceSyncPeriod == zero {
		obj.NamespaceSyncPeriod = metav1.Duration{Duration: 5 * time.Minute}
	}
}

func SetDefaults_NodeLifecycleControllerConfiguration(obj *NodeLifecycleControllerConfiguration) {
	zero := metav1.Duration{}
	if obj.PodEvictionTimeout == zero {
		obj.PodEvictionTimeout = metav1.Duration{Duration: 5 * time.Minute}
	}
	if obj.NodeMonitorGracePeriod == zero {
		obj.NodeMonitorGracePeriod = metav1.Duration{Duration: 40 * time.Second}
	}
	if obj.NodeStartupGracePeriod == zero {
		obj.NodeStartupGracePeriod = metav1.Duration{Duration: 60 * time.Second}
	}
	if obj.EnableTaintManager == nil {
		obj.EnableTaintManager = utilpointer.BoolPtr(true)
	}
}

func SetDefaults_ResourceQuotaControllerConfiguration(obj *ResourceQuotaControllerConfiguration) {
	zero := metav1.Duration{}
	if obj.ConcurrentResourceQuotaSyncs == 0 {
		obj.ConcurrentResourceQuotaSyncs = 5
	}
	if obj.ResourceQuotaSyncPeriod == zero {
		obj.ResourceQuotaSyncPeriod = metav1.Duration{Duration: 5 * time.Minute}
	}
}

func SetDefaults_PersistentVolumeRecyclerConfiguration(obj *PersistentVolumeRecyclerConfiguration) {
	if obj.MaximumRetry == 0 {
		obj.MaximumRetry = 3
	}
	if obj.MinimumTimeoutNFS == 0 {
		obj.MinimumTimeoutNFS = 300
	}
	if obj.IncrementTimeoutNFS == 0 {
		obj.IncrementTimeoutNFS = 30
	}
	if obj.MinimumTimeoutHostPath == 0 {
		obj.MinimumTimeoutHostPath = 60
	}
	if obj.IncrementTimeoutHostPath == 0 {
		obj.IncrementTimeoutHostPath = 30
	}
}

func SetDefaults_VolumeConfiguration(obj *VolumeConfiguration) {
	if obj.EnableHostPathProvisioning == nil {
		obj.EnableHostPathProvisioning = utilpointer.BoolPtr(false)
	}
	if obj.EnableDynamicProvisioning == nil {
		obj.EnableDynamicProvisioning = utilpointer.BoolPtr(true)
	}
	if obj.FlexVolumePluginDir == "" {
		obj.FlexVolumePluginDir = "/usr/libexec/kubernetes/kubelet-plugins/volume/exec/"
	}
}
