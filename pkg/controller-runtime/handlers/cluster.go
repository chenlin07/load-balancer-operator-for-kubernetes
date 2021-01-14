// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package handlers

import (
	akoov1alpha1 "gitlab.eng.vmware.com/core-build/ako-operator/api/v1alpha1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/cluster-api/util/conditions"
)

func SkipCluster(cluster *clusterv1.Cluster) bool {
	if cluster.Namespace == akoov1alpha1.TKGSystemNamespace {
		return true
	}
	if !conditions.IsTrue(cluster, clusterv1.ReadyCondition) {
		return true
	}
	return false
}