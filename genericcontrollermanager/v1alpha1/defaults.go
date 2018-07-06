/*
Copyright 2015 The Kubernetes Authors.

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

func SetDefaults_GenericComponentConfiguration(obj *GenericComponentConfiguration) {
	zero := metav1.Duration{}
	if obj.MinResyncPeriod == zero {
		obj.MinResyncPeriod = metav1.Duration{Duration: 12 * time.Hour}
	}
	if obj.ContentType == "" {
		obj.ContentType = "application/vnd.kubernetes.protobuf"
	}
	if obj.KubeAPIQPS == 0 {
		obj.KubeAPIQPS = 20.0
	}
	if obj.KubeAPIBurst == 0 {
		obj.KubeAPIBurst = 30
	}
	if obj.ControllerStartInterval == zero {
		obj.ControllerStartInterval = metav1.Duration{Duration: 0 * time.Second}
	}
}

func SetDefaults_KubeCloudSharedConfiguration(obj *KubeCloudSharedConfiguration) {
	zero := metav1.Duration{}
	// Port
	if obj.Address == "" {
		obj.Address = "0.0.0.0"
	}
	if obj.RouteReconciliationPeriod == zero {
		obj.RouteReconciliationPeriod = metav1.Duration{Duration: 10 * time.Second}
	}
	if obj.NodeMonitorPeriod == zero {
		obj.NodeMonitorPeriod = metav1.Duration{Duration: 5 * time.Second}
	}
	if obj.ClusterName == "" {
		obj.ClusterName = "kubernetes"
	}
	if obj.ConfigureCloudRoutes == nil {
		obj.ConfigureCloudRoutes = utilpointer.BoolPtr(true)
	}
}

func SetDefaults_ServiceControllerConfiguration(obj *ServiceControllerConfiguration) {
	if obj.ConcurrentServiceSyncs == 0 {
		obj.ConcurrentServiceSyncs = 1
	}
}
