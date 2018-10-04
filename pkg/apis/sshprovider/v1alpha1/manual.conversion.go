package v1alpha1

import (
	unsafe "unsafe"

	sshprovider "github.com/platform9/ssh-provider/pkg/apis/sshprovider"
	v1 "k8s.io/api/core/v1"
)

// Convert_sshprovider_ClusterSpec_To_v1alpha1_ClusterSpec converts unversioned type to
// v1alpha1 type. This has to manual as v1alpha1 does not have ClusterConfig
func Convert_sshprovider_ClusterSpec_To_v1alpha1_ClusterSpec(in *sshprovider.ClusterSpec, out *ClusterSpec) {
	//it cherry-picks the available fields and ignores ClusterConfig which is unavailable in v1alpha1
	out.EtcdCASecret = (*v1.LocalObjectReference)(unsafe.Pointer(in.EtcdCASecret))
	out.APIServerCASecret = (*v1.LocalObjectReference)(unsafe.Pointer(in.APIServerCASecret))
	out.FrontProxyCASecret = (*v1.LocalObjectReference)(unsafe.Pointer(in.FrontProxyCASecret))
	out.ServiceAccountKeySecret = (*v1.LocalObjectReference)(unsafe.Pointer(in.ServiceAccountKeySecret))
	out.BootstrapTokenSecret = (*v1.LocalObjectReference)(unsafe.Pointer(in.BootstrapTokenSecret))
	out.VIPConfiguration = (*VIPConfiguration)(unsafe.Pointer(in.VIPConfiguration))
}
