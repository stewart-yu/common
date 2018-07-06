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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/componentconfig"
)

// CloudProviderConfiguration contains basically elements about cloud provider.
type CloudProviderConfiguration struct {
	// Name is the provider for cloud services.
	Name string `json:"cloudProvider"`
	// cloudConfigFile is the path to the cloud provider configuration file.
	CloudConfigFile string `json:"cloudConfigFile"`
}

// DebuggingConfiguration contains elements used to debugging.
type DebuggingConfiguration struct {
	// enableProfiling enables profiling via web interface host:port/debug/pprof/
	EnableProfiling bool `json:"enableProfiling"`
	// EnableContentionProfiling enables lock contention profiling, if
	// EnableProfiling is true.
	EnableContentionProfiling bool `json:"enableContentionProfiling"`
}

// GenericComponentConfiguration contains generic elements shared by both kube-controller manager
// and cloud-controller manager.
type GenericComponentConfiguration struct {
	// minResyncPeriod is the resync period in reflectors; will be random between
	// minResyncPeriod and 2*minResyncPeriod.
	MinResyncPeriod metav1.Duration `json:"minResyncPeriod"`
	// contentType is contentType of requests sent to apiserver.
	ContentType string `json:"contentType"`
	// kubeAPIQPS is the QPS to use while talking with kubernetes apiserver.
	KubeAPIQPS float32 `json:"kubeAPIQPS"`
	// kubeAPIBurst is the burst to use while talking with kubernetes apiserver.
	KubeAPIBurst int32 `json:"kubeAPIBurst"`
	// How long to wait between starting controller managers
	ControllerStartInterval metav1.Duration `json:"controllerStartInterval"`
	// leaderElection defines the configuration of leader election client.
	LeaderElection componentconfig.LeaderElectionConfiguration `json:"leaderElection"`
}

// KubeCloudSharedConfiguration contains elements shared by both kube-controller manager
// and cloud-controller manager, but not generic.
type KubeCloudSharedConfiguration struct {
	// port is the port that the controller-manager's http service runs on.
	Port int32 `json:"port"`
	// address is the IP address to serve on (set to 0.0.0.0 for all interfaces).
	Address string `json:"address"`
	// useServiceAccountCredentials indicates whether controllers should be run with
	// individual service account credentials.
	UseServiceAccountCredentials bool `json:"useServiceAccountCredentials"`
	// run with untagged cloud instances
	AllowUntaggedCloud bool `json:"allowUntaggedCloud"`
	// routeReconciliationPeriod is the period for reconciling routes created for Nodes by cloud provider..
	RouteReconciliationPeriod metav1.Duration `json:"routeReconciliationPeriod"`
	// nodeMonitorPeriod is the period for syncing NodeStatus in NodeController.
	NodeMonitorPeriod metav1.Duration `json:"nodeMonitorPeriod"`
	// clusterName is the instance prefix for the cluster.
	ClusterName string `json:"clusterName"`
	// clusterCIDR is CIDR Range for Pods in cluster.
	ClusterCIDR string `json:"clusterCIDR"`
	// AllocateNodeCIDRs enables CIDRs for Pods to be allocated and, if
	// ConfigureCloudRoutes is true, to be set on the cloud provider.
	AllocateNodeCIDRs bool `json:"allocateNodeCIDRs"`
	// CIDRAllocatorType determines what kind of pod CIDR allocator will be used.
	CIDRAllocatorType string `json:"cIDRAllocatorType"`
	// configureCloudRoutes enables CIDRs allocated with allocateNodeCIDRs
	// to be configured on the cloud provider.
	ConfigureCloudRoutes *bool `json:"configureCloudRoutes"`
	// nodeSyncPeriod is the period for syncing nodes from cloudprovider. Longer
	// periods will result in fewer calls to cloud provider, but may delay addition
	// of new nodes to cluster.
	NodeSyncPeriod metav1.Duration `json:"nodeSyncPeriod"`
}

// ServiceControllerConfiguration contains elements describing ServiceController.
type ServiceControllerConfiguration struct {
	// concurrentServiceSyncs is the number of services that are
	// allowed to sync concurrently. Larger number = more responsive service
	// management, but more CPU (and network) load.
	ConcurrentServiceSyncs int32 `json:"concurrentServiceSyncs"`
}
